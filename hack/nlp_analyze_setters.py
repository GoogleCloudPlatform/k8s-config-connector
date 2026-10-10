#!/usr/bin/env python3

import json
import subprocess
import re
import sys

def extract_json(text):
    # Try to find a JSON code block
    match = re.search(r'```(?:json)?\n(.*?)\n```', text, re.DOTALL)
    if match:
        return match.group(1)
    
    # Try to find anything that looks like a JSON object
    match = re.search(r'\{.*\}', text, re.DOTALL)
    if match:
        return match.group(0)
    
    return text

def main():
    try:
        with open('gcp_mutability.json', 'r') as f:
            data = json.load(f)
    except Exception as e:
        print(f"Error reading gcp_mutability.json: {e}")
        sys.exit(1)

    total_resources = len(data)
    current = 0

    for kind, info in data.items():
        current += 1
        setters = info.get('setter_methods', {})
        
        if not setters:
            info['mutable_fields'] = []
            info['potentially_mutable_fields'] = []
            continue

        print(f"[{current}/{total_resources}] Analyzing {kind} with {len(setters)} setters...")
        
        prompt = f"""
Analyze the following setter methods for the GCP resource "{info.get('proto_message')}". 
Your goal is to classify which fields are definitively mutable, and which are potentially mutable (e.g., setting an external policy, or having an ambiguous description).

RULES:
1. For the field name, use 'matched_field_in_resource' if it is not empty. If it is empty, use 'inferred_field'.
2. If the description clearly states it changes/updates a property directly on the resource, it is 'mutable_fields'.
3. If it sets an external policy, a sub-resource, IAM policy, labels, or is otherwise ambiguous, it is 'potentially_mutable_fields'.
4. Output ONLY valid JSON, no markdown, no explanation.

Schema:
{{
  "mutable_fields": ["field_name_1"],
  "potentially_mutable_fields": ["field_name_2"]
}}

Setters Data:
{json.dumps(setters, indent=2)}
"""
        
        try:
            result = subprocess.run(
                ["gemini", "-p", prompt, "-o", "text"],
                capture_output=True,
                text=True,
                check=True
            )
            
            output = result.stdout.strip()
            clean_json = extract_json(output)
            
            try:
                parsed = json.loads(clean_json)
                info['mutable_fields'] = sorted(parsed.get('mutable_fields', []))
                info['potentially_mutable_fields'] = sorted(parsed.get('potentially_mutable_fields', []))
                print(f"  -> Mutable: {info['mutable_fields']}")
                print(f"  -> Potentially Mutable: {info['potentially_mutable_fields']}")
            except json.JSONDecodeError:
                print(f"  -> Failed to parse JSON from output:\n{output}")
                info['mutable_fields'] = []
                info['potentially_mutable_fields'] = []
                
        except subprocess.CalledProcessError as e:
            print(f"  -> Gemini CLI failed: {e.stderr}")
            info['mutable_fields'] = []
            info['potentially_mutable_fields'] = []

    try:
        with open('gcp_mutability.json', 'w') as f:
            json.dump(data, f, indent=2, sort_keys=True)
        print("\nSuccessfully updated gcp_mutability.json with NLP analysis.")
    except Exception as e:
        print(f"Error writing gcp_mutability.json: {e}")

if __name__ == "__main__":
    main()
