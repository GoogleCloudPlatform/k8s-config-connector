Simple dependency scanner tool

Keeps license detection easy - relies on a static set of files.

To generate in a nice format:

```
go run . scan --binary ~/bin/kpt | \
  jq -r '.[] | [.name, (.licenseInfo.licenseURLs | join(" ")), .licenseInfo.license, "kpt", "YES", if .licenseInfo.mustShipCode then  "YES" else "NO" end ] | @csv'
```

Saving a copy in the repo is helpful to track changes over time:

```
go run . scan --binary ~/bin/kpt | jq . > results.txt
```

To generate a LICENSES text file (useful for embedding):



```

go run . scan --print --binary ~/bin/kpt | jq -r '.[] | ("================================================================================\n= " + .name + " =\n\n" + .licenseFiles[].contents + "\n\n")' > ../../licenses/kpt.txt

```



## How to update license scan dependency



1. Run the `@dev/tasks/generate-licenses` script. This will create a directory called `modules` containing files marked `TODO`.

2. Run the following command to find and update the licenses:

   ```bash

   grep -r "TODO" | awk -F: '{print $1}' | xargs -I {} ./dev/tasks/find-license.sh {}

   ```

3. Move the updated modules to the correct location:

   ```bash

   cp -r -n modules/* experiments/tools/licensescan/modules/

   ```
