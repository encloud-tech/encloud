# encloud Desktop Application

The encloud Desktop Application is a cross-platform application (Windows, MacOS and Linux) built using the Wails framework. It combines a React based
frontend with the Golang based APIs into a seamless desktop experience for users who prefer a GUI. The Desktop App has 
an intuitive interface that is better suited for use by data analysts, data scientist and business analysts that prefer 
interacting with a UI.

The application can be downloaded for the relevant platform architecture via the [encloud website](https://encloud.tech/)

Check out the demo for the application and how to use

[![encloud Desktop App Demo](http://img.youtube.com/vi/VaXNkpykrPg/0.jpg)](https://www.youtube.com/watch?v=VaXNkpykrPg "encloud Desktop App Demo")

## Building the Desktop Application From Source

The Desktop application can be either downloaded for the website or built locally for the host architecture. This is 
especially useful for various Linux distros as compatibility of the Wails framework for various distros can vary, and we 
have tested our Linux build on Ubuntu 22.04. 

Make sure Wails CLI is installed, [read here](https://wails.io/docs/gettingstarted/installation)

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

### Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

### Building

To build a redistributable, production mode package, use `wails build`.
