import yaml 
import os


def read_yaml_file(crd_path):
    """
    Reads a YAML file from the specified path and returns its content as a Python dictionary.

    Args:
        crd_path (str): The full path to the YAML file.

    Returns:
        dict: A dictionary representing the content of the YAML file, or None if an error occurs.
    """
    if not os.path.exists(crd_path):
        raise(f"Error: File not found at '{crd_path}'")
    
    if not crd_path.lower().endswith(('.yaml', '.yml')):
        print(f"Warning: The file '{crd_path}' does not appear to be a YAML file.")

    try:
        with open(crd_path, 'r') as file:
            yaml_content = yaml.safe_load(file)
            return yaml_content
    except yaml.YAMLError as e:
        raise ValueError(f"Error parsing YAML file '{crd_path}': {e}")
    except Exception as e:
        raise Exception(f"An unexpected error occurred while reading '{crd_path}': {e}")

def write_yaml_file(file_path, content):
    """
    Writes the given content to a YAML file at the specified path.

    Args:
        file_path (str): The path to the file where content should be written.
        content (dict): The content to write to the file, as a Python dictionary.

    Returns:
        None: If the operation is successful.
        str: An error message if an exception occurs.
    """
    if not file_path:
        raise ("file_path cannot be None")
    if not content:
        raise ("content cannot be None")
    
    try:
        with open(file_path, 'w') as file:
            yaml.dump(content, file)
            return "File written successfully."
    except Exception as e:
        return f"An error occurred while writing to the file: {e}"  