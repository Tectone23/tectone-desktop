# Tectone SDK Development Documentation

### Overview

The

### Sections

	##### GO code
	The Go code sits at the root of the project and packages of the go code is in the `internal` following a Go convention where internally used packages are stored in an `internal` folder. This also serves us by keeping the folders at the root folder cleaner since it is also where other folders that have nothing to do with the go code reside.


	##### Frontend (Typescript/React)
	The `frontend` folder holds all the UI code which is written using HTML/CSS/Typescript/TSX. Inside the `src` folder the code is further divided in `assets`, `components`, `icons`, and `screens`. The names of the folders are self-explanatory but listed below are what they contain:

	assets - media files
	components - react components used within the UI
	icons - svg files of icons used in the UI
	screens - individual screens that make up the entire UI.


	##### Documentation
	Documentation folder contains the documentation pertaining to the development of the SDK GUI and therefore it is important to be updated regularly for when new features are added.


	##### Scripts
	Scripts folder holds the various scripts specific to the building of the code. The scripts are referenced in the `wails.json` configuration file and are called by wails before the build when the scripts are in the `pre-build` folder, and the `post-build` folder after the build is finished.

	##### Wireframes
	The wireframes folder holds the wireframes of the UI and are there for referencing, and they represent the wireframes for the different screens of the UI. When a new window is added to the UI, try to include the wireframe for it in this folder.
