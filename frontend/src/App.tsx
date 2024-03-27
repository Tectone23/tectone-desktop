import { HashRouter, Route, Routes } from "react-router-dom";
// import Home from "./screens/Home";
import Sidebar from "./components/Sidebar";
import NewProject from "./screens/NewProject";
import Logo from "./assets/images/White_Logo.webp";
import Home from "./screens/Home";
import Splash from "./screens/Splash";
import { useEffect, useState } from "react";
import Sandbox from "./screens/Sandbox";
import Containers from "./screens/Containers";
import Settings from "./screens/Settings";
import ProjectView from "./screens/ProjectView";
import ProjectSettings from "./screens/ProjectSettings";
import ErrorBanner from "./components/ErrorBanner";
import { useError, useInfo } from "./hooks/useMessaging";
import InfoBanner from "./components/InfoBanner";
import Feedback from "./screens/Feedback";
import FeedbackIframe from "./screens/FeedbackIframe";

function App() {
  const { error, setError } = useError();
  const { info, setInfo } = useInfo();
  const style = {
    draggable: {
      "--wails-draggable": "drag",
    } as React.CSSProperties,
    undraggable: {
      "--wails-draggable": "no-drag",
    } as React.CSSProperties,
  };

  return (
    <div className="mx-auto min-h-screen bg-gray-900  text-gray-400">
      <div
        className="fixed left-0 top-0 h-12 w-full items-center justify-between bg-gray-900 text-gray-400 z-50"
        style={style.draggable}
      >
        <div className="flex h-12 w-full pt-2 justify-end items-start">
          {/* <div className="h-6 w-[68px] bg-red-300" id="app-global-controls" />
          <div className="h-6 w-[68px] bg-blue-300" /> */}
          <div className="flex space-x-2 mr-4  text-gray-500 ">
            <button>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth="1.5"
                stroke="currentColor"
                className="w-6 h-6"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="m11.25 11.25.041-.02a.75.75 0 0 1 1.063.852l-.708 2.836a.75.75 0 0 0 1.063.853l.041-.021M21 12a9 9 0 1 1-18 0 9 9 0 0 1 18 0Zm-9-3.75h.008v.008H12V8.25Z"
                />
              </svg>
            </button>
            <button>
              <svg
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                strokeWidth="1.5"
                stroke="currentColor"
                className="w-6 h-6"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M12 6.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5ZM12 12.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5ZM12 18.75a.75.75 0 1 1 0-1.5.75.75 0 0 1 0 1.5Z"
                />
              </svg>
            </button>
          </div>
        </div>
      </div>
      <div
        style={style.undraggable}
        className="flex relative w-full h-screen max-h-screen max-w-full  overflow-x-hidden"
      >
        {error != "" ||
          (info != "" && (
            <div className="absolute w-[800px] px-20 top-8 z-50 left-[300px] ">
              {error != "" && (
                <ErrorBanner onAbort={() => setError("")} error={error} />
              )}

              {info != "" && (
                <InfoBanner onAbort={() => setInfo("")} info={info} />
              )}
            </div>
          ))}
        {/* <div className="absolute right-12 top-20 w-full">
          <ErrorBanner error="file could not be uploaded" />
        </div> */}
        <div className="flex h-full w-full">
          <HashRouter basename="/">
            <Routes>
              <Route path="/" element={<Splash />} />
              <Route path="/projects" element={<Home />} />
              <Route path="/newProject" element={<NewProject />} />
              <Route path="/sandbox" element={<Sandbox />} />
              <Route path="/containers" element={<Containers />} />
              <Route path="/settings" element={<Settings />} />
              <Route path="/feedback" element={<FeedbackIframe />} />

              <Route path="/project-view" element={<ProjectView />} />
              <Route path="/project-settings" element={<ProjectSettings />} />
            </Routes>
          </HashRouter>
        </div>
      </div>
    </div>
  );
}

export default App;
