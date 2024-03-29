# Jinja Expander docker image

Example usage:

(optional) Remove the existing expanded file.
```
rm -fr expanded/expanded
```

Run the expander on the `inputs/template` and `inputs/values`:
```
docker run --rm -v ${PWD}/tests/test1/inputs:/inputs -v ${PWD}/tests/test1/expanded:/expanded gcr.io/cdcs-test/expander-jinja2:latest /inputs/template /inputs/values --format=yaml -o expanded/expanded
```