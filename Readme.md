# Simple sitemap generator

## Run application

- clone repo 
- run `make compile` - binary will be created at ./bin/sitemap
- run `./bin/sitemap https://example.com`
- sitemap file `sitemap.xml` will be generated


### Additional arguments:
- parallel= number of parallel workers to navigate through site ( default: 4)
- output-file= output file path ( default: sitemap.xml )
- max-depth= max depth of url navigation recursion  ( default: 2 )

### Example usage

- `sitemap https://example.com -parallel=2 -max-depth=4 -output-file=example-com.xml`
