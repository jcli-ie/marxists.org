# Marxists.org

*I am not affiliated with marxists.org at this time.*

A reimaged front-end for marxists.org.

Data is scraped using the [marxists-parser](https://github.com/MarxistsDev/marxists-parser).

The UI is currently under heavy development, and I hope others more experienced with tailwind, htmx, and templ (templating for go) will join us.

Please feel free to share your ideas in the `Discussions` Tab.

[Watch the Demo](https://youtu.be/0b-LGEGfZuc)

## Development

Ensure you have Go 1.22 installed and your preferred version of Node.

_**Install Dependencies**_:

Install node dependencies via npm:

`$ npm install`

Install go dependencies:

`$ go mod download`

_**Hot Module Reloading**_:

If you want hot module reloading for easier development, install `Air` and `browser-sync`:

`$ go install air`

`$ npm install -g browser-sync`

Then run `npm run watch`, this will ensure the server is re-built if any of the Go or Templ files change.

If you want your browser to automatically update when working with templates, run `npm run sync` in a separate shell.

