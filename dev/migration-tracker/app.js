/**
 * Copyright 2026 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

let resourceData = [];
let downstreamMap = {};
let currentSort = { column: 'sortOrder', direction: 'asc' };
let expandedRows = new Set();
let activeTheme = 'light';

window.addEventListener('error', (event) => {
    const errorConsole = document.getElementById('error-console');
    const errorMessage = document.getElementById('error-message');
    if (errorConsole && errorMessage) {
        errorConsole.style.display = 'block';
        errorMessage.textContent = `${event.message} at ${event.filename}:${event.lineno}:${event.colno}`;
    }
});

window.addEventListener('unhandledrejection', (event) => {
    const errorConsole = document.getElementById('error-console');
    const errorMessage = document.getElementById('error-message');
    if (errorConsole && errorMessage) {
        errorConsole.style.display = 'block';
        errorMessage.textContent = `Unhandled Promise Rejection: ${event.reason}`;
    }
});
document.addEventListener('DOMContentLoaded', () => {
    // Theme configuration
    const savedTheme = localStorage.getItem('theme') || 'light';
    setTheme(savedTheme);

    fetchData();
});

function setTheme(theme) {
    activeTheme = theme;
    document.documentElement.setAttribute('data-theme', theme);
    localStorage.setItem('theme', theme);
    const themeBtn = document.getElementById('theme-btn');
    if (themeBtn) {
        themeBtn.innerHTML = theme === 'dark' 
            ? `<svg width="18" height="18" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path d="M12 3v1m0 16v1m9-9h-1M4 12H3m15.364-6.364l-.707.707M6.343 17.657l-.707.707m0-12.728l.707.707m11.314 11.314l.707.707M12 8a4 4 0 100 8 4 4 0 000-8z"></path></svg>`
            : `<svg width="18" height="18" fill="none" stroke="currentColor" stroke-width="2" viewBox="0 0 24 24"><path d="M20.354 15.354A9 9 0 018.646 3.646 9.003 9.003 0 0012 21a9.003 9.003 0 008.354-5.646z"></path></svg>`;
    }
}

function toggleTheme() {
    setTheme(activeTheme === 'light' ? 'dark' : 'light');
    if (document.getElementById('content-overview').classList.contains('active')) {
        renderCharts();
    }
    if (document.getElementById('content-progress') && document.getElementById('content-progress').classList.contains('active')) {
        renderProgressTab();
    }
}

function switchTab(tabId) {
    document.querySelectorAll('.menu-item').forEach(item => item.classList.remove('active'));
    document.querySelectorAll('.tab-content').forEach(content => content.classList.remove('active'));

    const menuMap = {
        overview: { tab: 'menu-overview', title: 'Dashboard Overview' },
        progress: { tab: 'menu-progress', title: 'Migration Progress & Projections' },
        matrix: { tab: 'menu-matrix', title: 'Resource Explorer' },
        explorer: { tab: 'menu-explorer', title: 'Dependency Topology' }
    };

    if (menuMap[tabId]) {
        document.getElementById(menuMap[tabId].tab).classList.add('active');
        document.getElementById(`content-${tabId}`).classList.add('active');
        document.getElementById('active-tab-title').textContent = menuMap[tabId].title;
        window.location.hash = tabId;
    }

    if (tabId === 'overview') {
        renderCharts();
    } else if (tabId === 'progress') {
        renderProgressTab();
    }
}

let migrationHistory = [
    { date: "2026-10-05", label: "Oct 05, '26", implemented: 80, flipped: 2 }
];

async function fetchMigrationHistory() {
    try {
        const response = await fetch('migration_history.json');
        if (response.ok) {
            const data = await response.json();
            if (Array.isArray(data) && data.length > 0) {
                migrationHistory = data;
            }
        }
    } catch (e) {
        console.warn('Using baseline migration history:', e);
    }
}

async function fetchData() {
    try {
        const [resData] = await Promise.all([
            fetch('data.json').then(r => r.json()),
            fetchMigrationHistory()
        ]);
        resourceData = resData;
        
        buildDownstreamMap();
        populateFilters();
        updateStats();
        renderCharts();
        renderGroupBreakdown();
        renderPriorityQueue();
        filterAndRenderTable();

        const hash = window.location.hash.substring(1);
        if (['overview', 'progress', 'matrix', 'explorer'].includes(hash)) {
            switchTab(hash);
        }
    } catch (error) {
        console.error('Error fetching data:', error);
        document.getElementById('matrix-tbody').innerHTML = `<tr><td colspan="9" style="text-align: center; color: var(--google-red);">Error loading data.json</td></tr>`;
    }
}

function buildDownstreamMap() {
    downstreamMap = {};
    resourceData.forEach(r => {
        r.dependencies.forEach(dep => {
            if (!downstreamMap[dep]) {
                downstreamMap[dep] = [];
            }
            downstreamMap[dep].push(r.kind);
        });
    });
}

function populateFilters() {
    const groups = [...new Set(resourceData.map(r => r.group))].sort();
    const groupSelect = document.getElementById('filter-group');
    const explorerSelect = document.getElementById('explorer-select');
    
    groupSelect.innerHTML = '<option value="All">All Groups</option>';
    explorerSelect.innerHTML = '<option value="">-- Choose a KRM Kind --</option>';

    groups.forEach(g => {
        groupSelect.innerHTML += `<option value="${g}">${g}</option>`;
    });

    const sortedKinds = [...resourceData].sort((a, b) => a.kind.localeCompare(b.kind));
    sortedKinds.forEach(r => {
        explorerSelect.innerHTML += `<option value="${r.kind}">${r.kind} (${r.group})</option>`;
    });
}

function getReadyToMigrate() {
    const completedSet = new Set(
        resourceData.filter(r => r.state === 'Completed').map(r => r.kind)
    );

    return resourceData.filter(r => {
        if (r.state === 'Completed') return false;
        if (r.edgeCases && r.edgeCases.gcpAPIDeprecated === true) return false;
        
        // Unblocked if all dependencies are in Completed state or not tracked in data.json
        return r.dependencies.every(dep => {
            const depResource = resourceData.find(item => item.kind === dep);
            if (!depResource) return true;
            return completedSet.has(dep);
        });
    });
}

function updateStats() {
    const activeResources = resourceData.filter(r => !(r.edgeCases && r.edgeCases.gcpAPIDeprecated === true));
    const total = activeResources.length;
    const completed = activeResources.filter(r => r.state === 'Completed').length;
    const inProgress = activeResources.filter(r => r.state === 'In Progress' || r.state === 'PR Sent').length;
    const readyCount = getReadyToMigrate().length;
    const adoptionRate = total > 0 ? Math.round((completed / total) * 100) : 0;

    document.getElementById('kpi-total').textContent = total;
    document.getElementById('kpi-migrated').textContent = completed;
    document.getElementById('kpi-progress').textContent = inProgress;
    document.getElementById('kpi-ready').textContent = readyCount;
    document.getElementById('kpi-adoption').textContent = `${adoptionRate}% adoption rate`;
    document.getElementById('ready-badge').textContent = `${readyCount} unblocked`;
}

function renderCharts() {
    const activeResources = resourceData.filter(r => !(r.edgeCases && r.edgeCases.gcpAPIDeprecated === true));
    const directCount = activeResources.filter(r => r.controllerType === 'Direct').length;
    const tfCount = activeResources.filter(r => r.controllerType === 'Terraform').length;
    const dclCount = activeResources.filter(r => r.controllerType === 'DCL').length;
    const totalCtrl = directCount + tfCount + dclCount;

    const directPct = totalCtrl > 0 ? Math.round((directCount / totalCtrl) * 100) : 0;
    const tfPct = totalCtrl > 0 ? Math.round((tfCount / totalCtrl) * 100) : 0;
    const dclPct = totalCtrl > 0 ? Math.round((dclCount / totalCtrl) * 100) : 0;

    // 1. Controller Donut Conic-Gradient
    if (totalCtrl > 0) {
        const directEnd = directPct;
        const tfEnd = directPct + tfPct;
        document.getElementById('donut-controller').style.background = 'conic-gradient(#1e8e3e 0% ' + directEnd + '%, #f9ab00 ' + directEnd + '% ' + tfEnd + '%, #1a73e8 ' + tfEnd + '% 100%)';
    } else {
        document.getElementById('donut-controller').style.background = 'conic-gradient(#dadce0 0% 100%)';
    }

    // Controller Legend
    document.getElementById('legend-controller').innerHTML = `
        <div class="legend-item">
            <div class="legend-label-group">
                <span class="legend-color-dot direct"></span>
                <span class="legend-name">Direct</span>
            </div>
            <span class="legend-value">${directCount} (${directPct}%)</span>
        </div>
        <div class="legend-item">
            <div class="legend-label-group">
                <span class="legend-color-dot terraform"></span>
                <span class="legend-name">Terraform</span>
            </div>
            <span class="legend-value">${tfCount} (${tfPct}%)</span>
        </div>
        <div class="legend-item">
            <div class="legend-label-group">
                <span class="legend-color-dot dcl"></span>
                <span class="legend-name">DCL</span>
            </div>
            <span class="legend-value">${dclCount} (${dclPct}%)</span>
        </div>
    `;

    // 2. States Donut Conic-Gradient
    const notStarted = activeResources.filter(r => r.state === 'Not Started').length;
    const inProgress = activeResources.filter(r => r.state === 'In Progress').length;
    const prSent = activeResources.filter(r => r.state === 'PR Sent').length;
    const completed = activeResources.filter(r => r.state === 'Completed').length;
    const totalStates = notStarted + inProgress + prSent + completed;

    const completedPct = totalStates > 0 ? Math.round((completed / totalStates) * 100) : 0;
    const prSentPct = totalStates > 0 ? Math.round((prSent / totalStates) * 100) : 0;
    const inProgressPct = totalStates > 0 ? Math.round((inProgress / totalStates) * 100) : 0;
    const notStartedPct = totalStates > 0 ? Math.round((notStarted / totalStates) * 100) : 0;

    if (totalStates > 0) {
        const completedEnd = completedPct;
        const prSentEnd = completedPct + prSentPct;
        const inProgressEnd = completedPct + prSentPct + inProgressPct;
        document.getElementById('donut-states').style.background = 'conic-gradient(#1e8e3e 0% ' + completedEnd + '%, #f9ab00 ' + completedEnd + '% ' + prSentEnd + '%, #8ab4f8 ' + prSentEnd + '% ' + inProgressEnd + '%, #5f6368 ' + inProgressEnd + '% 100%)';
    } else {
        document.getElementById('donut-states').style.background = 'conic-gradient(#dadce0 0% 100%)';
    }

    // States Legend
    document.getElementById('legend-states').innerHTML = `
        <div class="legend-item">
            <div class="legend-label-group">
                <span class="legend-color-dot completed"></span>
                <span class="legend-name">Completed</span>
            </div>
            <span class="legend-value">${completed} (${completedPct}%)</span>
        </div>
        <div class="legend-item">
            <div class="legend-label-group">
                <span class="legend-color-dot pr-sent"></span>
                <span class="legend-name">PR Sent</span>
            </div>
            <span class="legend-value">${prSent} (${prSentPct}%)</span>
        </div>
        <div class="legend-item">
            <div class="legend-label-group">
                <span class="legend-color-dot in-progress"></span>
                <span class="legend-name">In Progress</span>
            </div>
            <span class="legend-value">${inProgress} (${inProgressPct}%)</span>
        </div>
        <div class="legend-item">
            <div class="legend-label-group">
                <span class="legend-color-dot not-started"></span>
                <span class="legend-name">Not Started</span>
            </div>
            <span class="legend-value">${notStarted} (${notStartedPct}%)</span>
        </div>
    `;
}

function renderGroupBreakdown() {
    const stats = {};
    resourceData.forEach(r => {
        if (!stats[r.group]) {
            stats[r.group] = { total: 0, completed: 0 };
        }
        stats[r.group].total++;
        if (r.state === 'Completed') {
            stats[r.group].completed++;
        }
    });

    const sorted = Object.entries(stats)
        .map(([name, val]) => ({
            name,
            total: val.total,
            completed: val.completed,
            rate: Math.round((val.completed / val.total) * 100)
        }))
        .sort((a, b) => b.total - a.total);

    const list = document.getElementById('group-list');
    list.innerHTML = '';

    sorted.forEach(g => {
        list.innerHTML += `
            <div class="group-item">
                <div class="group-meta">
                    <span class="group-name">${g.name}</span>
                    <span>${g.completed}/${g.total} (${g.rate}%)</span>
                </div>
                <div class="group-bar-bg">
                    <div class="group-bar-fill" style="width: ${g.rate}%"></div>
                </div>
            </div>
        `;
    });
}

function renderPriorityQueue() {
    const ready = getReadyToMigrate();
    const list = document.getElementById('priority-list');
    list.innerHTML = '';

    if (ready.length === 0) {
        list.innerHTML = '<div style="text-align: center; color: var(--text-muted); padding: 32px 0;">All dependencies are currently migrated or blocked!</div>';
        return;
    }

    ready.sort((a, b) => a.sortOrder - b.sortOrder);

    ready.slice(0, 10).forEach(r => {
        list.innerHTML += `
            <div class="priority-item" onclick="selectInExplorer('${r.kind}')">
                <div class="priority-info">
                    <span class="priority-name">${r.kind}</span>
                    <span class="priority-group">${r.group} · Topo Order ${r.sortOrder}</span>
                </div>
                <span class="badge ${r.state === 'In Progress' ? 'in-progress' : 'not-started'}">${r.state}</span>
            </div>
        `;
    });
}

function filterAndRenderTable() {
    const search = document.getElementById('search-input').value.toLowerCase();
    const controller = document.getElementById('filter-controller').value;
    const state = document.getElementById('filter-state').value;
    const group = document.getElementById('filter-group').value;

    let filtered = resourceData.filter(r => {
        const matchesSearch = 
            r.kind.toLowerCase().includes(search) || 
            r.group.toLowerCase().includes(search) || 
            r.dependencies.some(d => d.toLowerCase().includes(search));
        
        const matchesController = controller === 'All' || r.controllerType === controller;
        const matchesState = state === 'All' || r.state === state;
        const matchesGroup = group === 'All' || r.group === group;

        return matchesSearch && matchesController && matchesState && matchesGroup;
    });

    // Apply sorting
    filtered.sort((a, b) => {
        let valA = a[currentSort.column];
        let valB = b[currentSort.column];

        if (typeof valA === 'string') valA = valA.toLowerCase();
        if (typeof valB === 'string') valB = valB.toLowerCase();

        if (valA < valB) return currentSort.direction === 'asc' ? -1 : 1;
        if (valA > valB) return currentSort.direction === 'asc' ? 1 : -1;
        return 0;
    });

    document.getElementById('meta-count').textContent = `Showing ${filtered.length} of ${resourceData.length} resources`;

    const tbody = document.getElementById('matrix-tbody');
    tbody.innerHTML = '';

    filtered.forEach(row => {
        const isExpanded = expandedRows.has(row.kind);
        const tr = document.createElement('tr');
        const stateClass = row.state.toLowerCase().replace(' ', '-');

        const stepDots = `
            <div class="step-dot-container">
                <span class="step-dot ${row.steps['gen-types'] ? 'done' : ''}" title="Gen Types"></span>
                <span class="step-dot ${row.steps['identity-reference'] ? 'done' : ''}" title="Identity Reference"></span>
                <span class="step-dot ${row.steps['mapper-fuzzer'] ? 'done' : ''}" title="Mapper Fuzzer"></span>
                <span class="step-dot ${row.steps.mocks ? 'done' : ''}" title="Mocks"></span>
                <span class="step-dot ${row.steps.controller ? 'done' : ''}" title="Controller"></span>
                <span class="step-dot ${row.steps.tests ? 'done' : ''}" title="Tests"></span>
            </div>
        `;

        const depsHtml = row.dependencies.map(d => {
            const depR = resourceData.find(item => item.kind === d);
            const isBlocked = depR && depR.state !== 'Completed';
            return `<span class="dep-tag ${isBlocked ? 'blocking' : ''}" onclick="selectInExplorer('${d}')">${d}</span>`;
        }).join('') || '-';

        const isDeprecated = row.edgeCases && row.edgeCases.gcpAPIDeprecated === true;
        const kindDisplay = isDeprecated ? `<span style="text-decoration: line-through; color: var(--text-muted); font-style: italic;">${row.kind}</span>` : `<strong>${row.kind}</strong>`;
        const deprecationBadge = isDeprecated ? `<span class="badge deprecated" title="${row.edgeCases.notes || ''}">GCP API Deprecated</span>` : '';

        tr.innerHTML = `
            <td>
                <button class="detail-toggle" onclick="toggleRowDetails('${row.kind}')">
                    ${isExpanded ? '▼' : '▶'}
                </button>
            </td>
            <td>${kindDisplay}</td>
            <td>${row.group}</td>
            <td>${row.sortOrder === 9999 ? '-' : row.sortOrder}</td>
            <td>${row.downstreamCount}</td>
            <td>
                <div style="font-weight: 500;">${row.defaultController}</div>
                <div style="font-size: 11px; color: var(--text-muted);">Supported: ${row.supportedControllers.join(', ')}</div>
            </td>
            <td>
                <span class="badge ${stateClass}">${row.state}</span>
                ${deprecationBadge}
            </td>
            <td>${stepDots}</td>
            <td><div class="dependency-tags">${depsHtml}</div></td>
        `;
        tbody.appendChild(tr);

        if (isExpanded) {
            const detailsTr = document.createElement('tr');
            detailsTr.className = 'row-expanded-details';

            const downstream = downstreamMap[row.kind] || [];
            const downstreamTags = downstream.map(d => `<span class="dep-tag" onclick="selectInExplorer('${d}')">${d}</span>`).join('') || 'None';

            let notesHtml = '';
            if (row.notes) {
                notesHtml = `<p><strong>Notes:</strong> <span>${row.notes}</span></p>`;
            }
            let edgeCaseNoteHtml = '';
            if (row.edgeCases && row.edgeCases.notes) {
                edgeCaseNoteHtml = `<p><strong>GCP API Deprecation Info:</strong> <span style="color: var(--google-red); font-weight: 500;">${row.edgeCases.notes}</span></p>`;
            }
            if (!row.notes && !edgeCaseNoteHtml) {
                notesHtml = `<p><strong>Notes:</strong> <span style="color: var(--text-muted);">None</span></p>`;
            }

            detailsTr.innerHTML = `
                <td></td>
                <td colspan="8">
                    <div class="drawer-content">
                        <div class="drawer-column">
                            <h4>Migration Notes</h4>
                            <p><strong>Mocks Last Refreshed:</strong> ${row.mocksLastRefreshed || 'Never'}</p>
                            ${notesHtml}
                            ${edgeCaseNoteHtml}
                        </div>
                        <div class="drawer-column">
                            <h4>Blocked Downstream Elements</h4>
                            <div class="dependency-tags">${downstreamTags}</div>
                        </div>
                    </div>
                </td>
            `;
            tbody.appendChild(detailsTr);
        }
    });
}

// Staging and sorting details
function toggleRowDetails(kind) {
    if (expandedRows.has(kind)) {
        expandedRows.delete(kind);
    } else {
        expandedRows.add(kind);
    }
    filterAndRenderTable();
}

function sortExplorer(col) {
    if (currentSort.column === col) {
        currentSort.direction = currentSort.direction === 'asc' ? 'desc' : 'asc';
    } else {
        currentSort.column = col;
        currentSort.direction = 'asc';
    }

    document.querySelectorAll('thead th span').forEach(s => s.textContent = '');
    const activeSpan = document.getElementById(`sort-${col}`);
    if (activeSpan) {
        activeSpan.textContent = currentSort.direction === 'asc' ? ' ▲' : ' ▼';
    }

    filterAndRenderTable();
}

function selectInExplorer(kind) {
    switchTab('explorer');
    const select = document.getElementById('explorer-select');
    select.value = kind;
    loadDependencyTree();
}

function loadDependencyTree() {
    const kind = document.getElementById('explorer-select').value;
    
    const targetBox = document.getElementById('explorer-target-box');
    const upstreamList = document.getElementById('explorer-upstream-list');
    const downstreamList = document.getElementById('explorer-downstream-list');

    targetBox.innerHTML = '';
    upstreamList.innerHTML = '';
    downstreamList.innerHTML = '';

    if (!kind) {
        targetBox.innerHTML = '<div style="text-align: center; color: var(--text-muted); padding: 24px;">Please select a resource kind above.</div>';
        return;
    }

    const r = resourceData.find(item => item.kind === kind);
    if (!r) return;

    // Target Node
    targetBox.innerHTML = `
        <div class="node-box active">
            <div>
                <strong>${r.kind}</strong>
                <div style="font-size: 11px; color: var(--text-secondary); margin-top: 4px;">Topo Order: ${r.sortOrder}</div>
            </div>
            <span class="badge ${r.state.toLowerCase().replace(' ', '-')}">${r.state}</span>
        </div>
    `;

    // Upstream Dependencies
    if (r.dependencies.length === 0) {
        upstreamList.innerHTML = '<div style="color: var(--text-muted); font-size: 13px; text-align: center; padding: 24px 0;">No dependencies.</div>';
    } else {
        r.dependencies.forEach(dep => {
            const depR = resourceData.find(item => item.kind === dep);
            const state = depR ? depR.state : 'Unknown';
            const stateClass = state.toLowerCase().replace(' ', '-');
            upstreamList.innerHTML += `
                <div class="node-box" onclick="selectInExplorer('${dep}')">
                    <span>${dep}</span>
                    <span class="badge ${stateClass}">${state}</span>
                </div>
            `;
        });
    }

    // Downstream Blocks
    const downstream = downstreamMap[kind] || [];
    if (downstream.length === 0) {
        downstreamList.innerHTML = '<div style="color: var(--text-muted); font-size: 13px; text-align: center; padding: 24px 0;">Does not block other resources.</div>';
    } else {
        downstream.forEach(down => {
            const downR = resourceData.find(item => item.kind === down);
            const state = downR ? downR.state : 'Unknown';
            const stateClass = state.toLowerCase().replace(' ', '-');
            downstreamList.innerHTML += `
                <div class="node-box" onclick="selectInExplorer('${down}')">
                    <span>${down}</span>
                    <span class="badge ${stateClass}">${state}</span>
                </div>
            `;
        });
    }
}

function renderProgressTab() {
    const container = document.getElementById('progress-svg-wrapper');
    if (!container) return;

    const startDate = new Date('2026-10-05T00:00:00Z');
    const targetDate = new Date('2027-03-31T00:00:00Z');
    const totalDays = (targetDate - startDate) / (1000 * 60 * 60 * 24); // 177 days
    const totalWeeks = totalDays / 7.0; // ~25.2857 weeks
    const totalScope = 301;
    const startImpl = 80;
    const startFlip = 2;

    // SVG geometry
    const svgWidth = 1000;
    const svgHeight = 320;
    const padLeft = 40;
    const padRight = 70; // 930px line end, leaves room for Y-axis numbers at 940px
    const padTop = 30;
    const padBottom = 45; // 275px bottom line

    const plotWidth = (svgWidth - padRight) - padLeft; // 890px
    const plotHeight = (svgHeight - padBottom) - padTop; // 245px

    function getX(weekFrac) {
        return padLeft + (weekFrac / totalWeeks) * plotWidth;
    }

    function getY(val) {
        return (svgHeight - padBottom) - (val / totalScope) * plotHeight;
    }

    // Grid lines for Y (0, 50, 100, 150, 200, 250, 301)
    const yTicks = [0, 50, 100, 150, 200, 250, 301];
    let gridLinesHtml = '';
    const isDark = activeTheme === 'dark';
    const gridStroke = isDark ? 'rgba(255, 255, 255, 0.12)' : 'rgba(0, 0, 0, 0.1)';
    const textColor = isDark ? '#9aa0a6' : '#5f6368';
    const textPrimary = isDark ? '#e8eaed' : '#202124';

    yTicks.forEach(val => {
        const y = getY(val).toFixed(1);
        gridLinesHtml += `<line x1="${padLeft}" y1="${y}" x2="${svgWidth - padRight}" y2="${y}" stroke="${gridStroke}" stroke-width="1" />`;
        gridLinesHtml += `<text x="${svgWidth - padRight + 10}" y="${(parseFloat(y) + 4).toFixed(1)}" fill="${textColor}" font-size="10" text-anchor="start">${val}</text>`;
    });

    // Time points & X-ticks
    const numIntervals = 26;
    const weekPoints = [];
    for (let i = 0; i <= numIntervals; i++) {
        const curDate = new Date(startDate.getTime() + i * 7 * 24 * 60 * 60 * 1000);
        if (curDate > targetDate) break;
        const weekFrac = (curDate - startDate) / (1000 * 60 * 60 * 24 * 7);
        const linImpl = startImpl + (weekFrac / totalWeeks) * (totalScope - startImpl);
        const linFlip = startFlip + (weekFrac / totalWeeks) * (totalScope - startFlip);
        const label = curDate.toLocaleDateString('en-US', { month: 'short', day: '2-digit', year: '2-digit' }).replace(',', " '");
        weekPoints.push({
            index: i,
            weekFrac: weekFrac,
            date: curDate,
            label: label,
            linImpl: linImpl,
            linFlip: linFlip,
            x: getX(weekFrac)
        });
    }

    // Add target point
    const lastPt = weekPoints[weekPoints.length - 1];
    if (lastPt && lastPt.weekFrac < totalWeeks) {
        weekPoints.push({
            index: numIntervals + 1,
            weekFrac: totalWeeks,
            date: targetDate,
            label: "Mar 31, '27",
            linImpl: totalScope,
            linFlip: totalScope,
            x: getX(totalWeeks)
        });
    }

    // X Axis ticks every 3 intervals
    let xTicksHtml = '';
    weekPoints.forEach((pt, idx) => {
        if (idx % 3 === 0 || idx === weekPoints.length - 1) {
            xTicksHtml += `<text x="${pt.x.toFixed(1)}" y="${(svgHeight - padBottom + 22).toFixed(1)}" fill="${textColor}" font-size="10" text-anchor="middle">${pt.label}</text>`;
        }
    });

    // Projected Paths
    const projImplPath = weekPoints.map((pt, idx) => `${idx === 0 ? 'M' : 'L'} ${pt.x.toFixed(1)} ${getY(pt.linImpl).toFixed(1)}`).join(' ');
    const projFlipPath = weekPoints.map((pt, idx) => `${idx === 0 ? 'M' : 'L'} ${pt.x.toFixed(1)} ${getY(pt.linFlip).toFixed(1)}`).join(' ');

    // Actual Data Points
    const actualPoints = migrationHistory.map(item => {
        const itemDate = new Date(item.date + 'T00:00:00Z');
        const weekFrac = Math.max(0, (itemDate - startDate) / (1000 * 60 * 60 * 24 * 7));
        const x = getX(weekFrac);
        const yImpl = getY(item.implemented);
        const yFlip = getY(item.flipped);

        const projImpl = startImpl + (weekFrac / totalWeeks) * (totalScope - startImpl);
        const projFlip = startFlip + (weekFrac / totalWeeks) * (totalScope - startFlip);

        return {
            dateStr: item.date,
            label: item.label,
            weekFrac: weekFrac,
            implemented: item.implemented,
            flipped: item.flipped,
            projImpl: projImpl,
            projFlip: projFlip,
            implDelta: item.implemented - projImpl,
            flipDelta: item.flipped - projFlip,
            x: x,
            yImpl: yImpl,
            yFlip: yFlip
        };
    });

    const realImplPath = actualPoints.map((pt, idx) => `${idx === 0 ? 'M' : 'L'} ${pt.x.toFixed(1)} ${pt.yImpl.toFixed(1)}`).join(' ');
    const realFlipPath = actualPoints.map((pt, idx) => `${idx === 0 ? 'M' : 'L'} ${pt.x.toFixed(1)} ${pt.yFlip.toFixed(1)}`).join(' ');

    let realMarkersHtml = '';
    actualPoints.forEach(pt => {
        realMarkersHtml += `<circle cx="${pt.x.toFixed(1)}" cy="${pt.yImpl.toFixed(1)}" r="4.5" fill="#10b981" stroke="var(--bg-card)" stroke-width="2" />`;
        realMarkersHtml += `<circle cx="${pt.x.toFixed(1)}" cy="${pt.yFlip.toFixed(1)}" r="4.5" fill="#06b6d4" stroke="var(--bg-card)" stroke-width="2" />`;
    });

    const latestActual = actualPoints[actualPoints.length - 1] || actualPoints[0];

    // Update legend actual numbers
    const legendRealImpl = document.getElementById('legend-real-impl');
    const legendRealFlip = document.getElementById('legend-real-flip');
    if (legendRealImpl) legendRealImpl.textContent = latestActual.implemented;
    if (legendRealFlip) legendRealFlip.textContent = latestActual.flipped;

    // Advisory cards evaluation
    updateAdvisories(latestActual);

    // Build SVG
    const svgHtml = `
    <svg id="svgProgressChart" viewBox="0 0 ${svgWidth} ${svgHeight}" width="100%" height="100%" style="overflow: visible; font-family: var(--font-family);">
      <!-- Grid -->
      ${gridLinesHtml}
      ${xTicksHtml}

      <!-- Right Axis Header Only -->
      <text x="${svgWidth - padRight}" y="18" fill="${textPrimary}" font-size="11" font-weight="600" text-anchor="end">Controllers (0 &ndash; 301)</text>

      <!-- Dotted Projected Lines -->
      <path d="${projImplPath}" fill="none" stroke="#34d399" stroke-width="2.5" stroke-dasharray="4,3" opacity="0.9" />
      <path d="${projFlipPath}" fill="none" stroke="#22d3ee" stroke-width="2.5" stroke-dasharray="4,3" opacity="0.9" />

      <!-- Solid Real Progress Lines -->
      <path d="${realImplPath}" fill="none" stroke="#10b981" stroke-width="3" />
      <path d="${realFlipPath}" fill="none" stroke="#06b6d4" stroke-width="3" />

      <!-- Real Markers -->
      ${realMarkersHtml}

      <!-- Vertical Guideline & Tooltip Indicator -->
      <line id="prog-hover-line" x1="0" y1="${padTop}" x2="0" y2="${svgHeight - padBottom}" stroke="${textColor}" stroke-width="1.5" stroke-dasharray="2,2" style="display: none;" />
      <circle id="prog-hover-impl" cx="0" cy="0" r="5" fill="#10b981" stroke="var(--bg-card)" stroke-width="2" style="display: none;" />
      <circle id="prog-hover-flip" cx="0" cy="0" r="5" fill="#06b6d4" stroke="var(--bg-card)" stroke-width="2" style="display: none;" />
    </svg>
    <div id="prog-tooltip" style="display: none; position: absolute; background: rgba(15, 23, 42, 0.92); color: #f8fafc; padding: 10px 14px; border-radius: 6px; font-size: 11px; pointer-events: none; border: 1px solid rgba(255,255,255,0.15); box-shadow: 0 4px 12px rgba(0,0,0,0.3); z-index: 100; min-width: 170px;"></div>
    `;

    container.innerHTML = svgHtml;

    // Attach mouse interaction
    const svgElem = document.getElementById('svgProgressChart');
    const hoverLine = document.getElementById('prog-hover-line');
    const hoverImpl = document.getElementById('prog-hover-impl');
    const hoverFlip = document.getElementById('prog-hover-flip');
    const tooltip = document.getElementById('prog-tooltip');

    svgElem.addEventListener('mousemove', (e) => {
        const rect = svgElem.getBoundingClientRect();
        const mouseX = e.clientX - rect.left;
        const svgMouseX = (mouseX / rect.width) * svgWidth;

        if (svgMouseX < padLeft || svgMouseX > (svgWidth - padRight)) {
            hoverLine.style.display = 'none';
            hoverImpl.style.display = 'none';
            hoverFlip.style.display = 'none';
            tooltip.style.display = 'none';
            return;
        }

        const weekFrac = ((svgMouseX - padLeft) / plotWidth) * totalWeeks;
        const curDate = new Date(startDate.getTime() + weekFrac * 7 * 24 * 60 * 60 * 1000);
        const curLabel = curDate.toLocaleDateString('en-US', { month: 'short', day: '2-digit', year: '2-digit' }).replace(',', " '");

        const curLinImpl = Math.min(totalScope, startImpl + (weekFrac / totalWeeks) * (totalScope - startImpl));
        const curLinFlip = Math.min(totalScope, startFlip + (weekFrac / totalWeeks) * (totalScope - startFlip));

        // Find if actual point exists near this week
        const matchActual = actualPoints.find(p => Math.abs(p.weekFrac - weekFrac) < 0.5);

        const xPos = svgMouseX;
        const yLin = getY(curLinImpl);
        const yFlip = getY(curLinFlip);

        hoverLine.setAttribute('x1', xPos);
        hoverLine.setAttribute('x2', xPos);
        hoverLine.style.display = 'block';

        hoverImpl.setAttribute('cx', xPos);
        hoverImpl.setAttribute('cy', yLin);
        hoverImpl.style.display = 'block';

        hoverFlip.setAttribute('cx', xPos);
        hoverFlip.setAttribute('cy', yFlip);
        hoverFlip.style.display = 'block';

        let html = `<div style="font-weight: 600; margin-bottom: 6px; border-bottom: 1px solid rgba(255,255,255,0.2); padding-bottom: 4px;">${curLabel} (W${weekFrac.toFixed(1)})</div>`;
        if (matchActual) {
            html += `<div style="color: #10b981; font-weight: 600;">Real Impl: ${matchActual.implemented} <span style="font-size: 10px; color: ${matchActual.implDelta >= 0 ? '#34d399' : '#f87171'};">(${matchActual.implDelta >= 0 ? '+' : ''}${matchActual.implDelta.toFixed(1)})</span></div>`;
            html += `<div style="color: #06b6d4; font-weight: 600;">Real Flip: ${matchActual.flipped} <span style="font-size: 10px; color: ${matchActual.flipDelta >= 0 ? '#22d3ee' : '#f87171'};">(${matchActual.flipDelta >= 0 ? '+' : ''}${matchActual.flipDelta.toFixed(1)})</span></div>`;
        }
        html += `<div style="color: #34d399; margin-top: 4px;">Target Impl: ${curLinImpl.toFixed(1)}</div>`;
        html += `<div style="color: #22d3ee;">Target Flip: ${curLinFlip.toFixed(1)}</div>`;

        tooltip.innerHTML = html;
        tooltip.style.display = 'block';

        let leftPx = mouseX + 15;
        if (leftPx + 180 > rect.width) {
            leftPx = mouseX - 190;
        }
        tooltip.style.left = `${leftPx}px`;
        tooltip.style.top = `${Math.min(rect.height - 120, Math.max(10, (e.clientY - rect.top) - 40))}px`;
    });

    svgElem.addEventListener('mouseleave', () => {
        hoverLine.style.display = 'none';
        hoverImpl.style.display = 'none';
        hoverFlip.style.display = 'none';
        tooltip.style.display = 'none';
    });
}

function updateAdvisories(latestActual) {
    const implCard = document.getElementById('advisory-impl-card');
    const implIcon = document.getElementById('advisory-impl-icon');
    const implTitle = document.getElementById('advisory-impl-title');
    const implDesc = document.getElementById('advisory-impl-desc');

    const flipCard = document.getElementById('advisory-flip-card');
    const flipIcon = document.getElementById('advisory-flip-icon');
    const flipTitle = document.getElementById('advisory-flip-title');
    const flipDesc = document.getElementById('advisory-flip-desc');

    if (!implCard || !flipCard) return;

    implCard.classList.remove('ahead', 'ontime', 'behind');
    if (latestActual.implDelta > 0.5) {
        implCard.classList.add('ahead');
        implIcon.textContent = '🚀';
        implTitle.textContent = `Implementation: Ahead of Schedule (+${latestActual.implDelta.toFixed(1)} controllers)`;
        implDesc.textContent = `Actual: ${latestActual.implemented} vs Target: ${latestActual.projImpl.toFixed(1)}. Pacing ahead of required 8.7/week run rate!`;
    } else if (latestActual.implDelta < -0.5) {
        implCard.classList.add('behind');
        implIcon.textContent = '⚠️';
        implTitle.textContent = `Implementation: Behind Schedule (${latestActual.implDelta.toFixed(1)} controllers)`;
        implDesc.textContent = `Actual: ${latestActual.implemented} vs Target: ${latestActual.projImpl.toFixed(1)}. Velocity must increase to meet March 31, 2027 target!`;
    } else {
        implCard.classList.add('ontime');
        implIcon.textContent = '✅';
        implTitle.textContent = `Implementation: On Track (Current: ${latestActual.implemented} / Target: ${latestActual.projImpl.toFixed(1)})`;
        implDesc.textContent = `Tracking required linear delivery pace (8.7 controllers/week). 221 remaining to implement.`;
    }

    flipCard.classList.remove('ahead', 'ontime', 'behind');
    if (latestActual.flipDelta > 0.5) {
        flipCard.classList.add('ahead');
        flipIcon.textContent = '🚀';
        flipTitle.textContent = `Default Flip: Ahead of Schedule (+${latestActual.flipDelta.toFixed(1)} flips)`;
        flipDesc.textContent = `Actual: ${latestActual.flipped} vs Target: ${latestActual.projFlip.toFixed(1)}. Pacing ahead of required 11.8/week run rate!`;
    } else if (latestActual.flipDelta < -0.5) {
        flipCard.classList.add('behind');
        flipIcon.textContent = '⚠️';
        flipTitle.textContent = `Default Flip: Behind Schedule (${latestActual.flipDelta.toFixed(1)} flips)`;
        flipDesc.textContent = `Actual: ${latestActual.flipped} vs Target: ${latestActual.projFlip.toFixed(1)}. Qualification and default flipping cadence needs acceleration.`;
    } else {
        flipCard.classList.add('ontime');
        flipIcon.textContent = '✅';
        flipTitle.textContent = `Default Flip: On Track (Current: ${latestActual.flipped} / Target: ${latestActual.projFlip.toFixed(1)})`;
        flipDesc.textContent = `Tracking required linear delivery pace (11.8 flips/week). 299 flips needed.`;
    }
}

