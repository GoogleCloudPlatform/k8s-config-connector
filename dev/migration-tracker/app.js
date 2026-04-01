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
let currentSort = { column: 'sortOrder', direction: 'asc' };

document.addEventListener('DOMContentLoaded', () => {
    fetchData();

    document.getElementById('search').addEventListener('input', renderTable);
    document.getElementById('filter-state').addEventListener('change', renderTable);

    document.querySelectorAll('th[data-sort]').forEach(th => {
        th.addEventListener('click', () => {
            const column = th.getAttribute('data-sort');
            if (currentSort.column === column) {
                currentSort.direction = currentSort.direction === 'asc' ? 'desc' : 'asc';
            } else {
                currentSort.column = column;
                currentSort.direction = 'asc';
            }
            updateSortIcons();
            renderTable();
        });
    });
});

async function fetchData() {
    try {
        const response = await fetch('data.json');
        resourceData = await response.json();
        updateStats();
        updateSortIcons();
        renderTable();
    } catch (error) {
        console.error('Error fetching data:', error);
        document.getElementById('table-body').innerHTML = `<tr><td colspan="11" style="text-align: center; color: red;">Error loading data.json</td></tr>`;
    }
}

function updateStats() {
    const total = resourceData.length;
    const completed = resourceData.filter(r => r.state === 'Completed').length;
    const inProgress = resourceData.filter(r => r.state === 'In Progress').length;

    document.getElementById('stat-total').textContent = total;
    document.getElementById('stat-completed').textContent = completed;
    document.getElementById('stat-inprogress').textContent = inProgress;
}

function updateSortIcons() {
    document.querySelectorAll('th[data-sort] .sort-icon').forEach(icon => icon.textContent = '');
    const activeTh = document.querySelector(`th[data-sort="${currentSort.column}"] .sort-icon`);
    if (activeTh) {
        activeTh.textContent = currentSort.direction === 'asc' ? ' ▲' : ' ▼';
    }
}

function renderTable() {
    const searchVal = document.getElementById('search').value.toLowerCase();
    const stateVal = document.getElementById('filter-state').value;

    let filteredData = resourceData.filter(row => {
        const matchesSearch = 
            row.kind.toLowerCase().includes(searchVal) || 
            row.group.toLowerCase().includes(searchVal) ||
            row.dependencies.some(d => d.toLowerCase().includes(searchVal));
        
        const matchesState = stateVal === 'All' || row.state === stateVal;

        return matchesSearch && matchesState;
    });

    filteredData.sort((a, b) => {
        let valA = a[currentSort.column];
        let valB = b[currentSort.column];

        if (typeof valA === 'string') valA = valA.toLowerCase();
        if (typeof valB === 'string') valB = valB.toLowerCase();

        if (valA < valB) return currentSort.direction === 'asc' ? -1 : 1;
        if (valA > valB) return currentSort.direction === 'asc' ? 1 : -1;
        return 0;
    });

    const tbody = document.getElementById('table-body');
    tbody.innerHTML = '';

    filteredData.forEach(row => {
        const tr = document.createElement('tr');

        const stateClass = row.state.toLowerCase().replace(' ', '-');
        
        // Steps mapping
        const stepsHtml = `
            <span class="step-indicator ${row.steps['gen-types'] ? 'done' : ''}" title="Gen Types"></span>
            <span class="step-indicator ${row.steps['identity-reference'] ? 'done' : ''}" title="Identity Reference"></span>
            <span class="step-indicator ${row.steps['mapper-fuzzer'] ? 'done' : ''}" title="Mapper Fuzzer"></span>
            <span class="step-indicator ${row.steps.mocks ? 'done' : ''}" title="Mocks"></span>
            <span class="step-indicator ${row.steps.controller ? 'done' : ''}" title="Controller"></span>
            <span class="step-indicator ${row.steps.tests ? 'done' : ''}" title="Tests"></span>
        `;

        const depsHtml = row.dependencies.map(d => `<span class="dependency-tag">${d}</span>`).join('') || '-';

        tr.innerHTML = `
            <td><strong>${row.kind}</strong></td>
            <td>${row.group}</td>
            <td>${row.version}</td>
            <td>${row.sortOrder === 9999 ? '-' : row.sortOrder}</td>
            <td>${row.downstreamCount}</td>
            <td>
                <div style="font-weight: 600;">${row.defaultController}</div>
                <div style="font-size: 11px; color: #5f6368;">Supported: ${row.supportedControllers.join(', ')}</div>
            </td>
            <td><span class="badge ${stateClass}">${row.state}</span></td>
            <td>${stepsHtml}</td>
            <td>${depsHtml}</td>
            <td>${row.mocksLastRefreshed}</td>
            <td style="color: #d93025; font-size: 11px;">${row.notes || ''}</td>
        `;
        tbody.appendChild(tr);
    });
}
