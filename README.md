# click-to-cloud

WARNING: This is no longer maintained. I recommend the [Heroku Button](https://devcenter.heroku.com/articles/heroku-button) as an excellent alternative.

Make it possible to deploy applications to heroku and other places with just a click. 

## Usage

```bash
click-to-cloud --repo https://github.com/scottmotte/handshake.git
```

Replace the repo with the repo you want to deploy. It needs to be a click-to-cloud enabled repo. It it is
it will have a a `click-to-cloud.json` file in the root of its project directory.

## How does it work

The developer defines a set of deployment rules in their application inside a click-to-cloud.json file. 
The click-to-cloud script takes care of the rest - knowing only the location of the git repo.

## Installation

You need to install the click-to-cloud binary on your machine. Do the following to do so. You must have go.

```bash
git clone https://github.com/scottmotte/click-to-cloud.git
cd click-to-cloud
go install
```

