import { BrowserOpenURL } from "../../wailsjs/runtime/runtime";
import Sidebar from "./Sidebar";
import Logo from "../assets/images/White_Logo.webp";
import { NavLink } from "react-router-dom";

interface ScreenProps {
  title: string;
}

export default function Screen() {
  return (
    <>
      <div className="w-[300px] max-h-full  border-r border-gray-800 pt-8  bg-gray-900 h-full px-6">
        <div className="flex justify-between items-center  pb-12 border-b border-gray-700 mb-12">
          <img src={Logo} alt="" />
        </div>
        <Sidebar />
      </div>
      <div className="flex w-full max-h-full overflow-y-scroll text-white px-4">
        <div className="flex flex-col w-full mt-20">
          <div className="mt-4 mb-6  text-sm text-purple-400 underline flex space-x-4">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth={1.5}
              stroke="currentColor"
              className="w-6 h-6"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M17.25 8.25 21 12m0 0-3.75 3.75M21 12H3"
              />
            </svg>

            <button
              onClick={() => {
                BrowserOpenURL("https://developer.algorand.org/docs/");
              }}
            >
              Algorand Developer Documentation
            </button>
          </div>
          <div className="sm:flex sm:items-center">
            <div className="sm:flex-auto">
              <h1 className="text-base font-semibold leading-6 text-white">
                Projects
              </h1>
              <p className="mt-2 text-sm text-gray-300">
                A list of all the projects in your development environment.
              </p>
            </div>
            <div className="mt-4 sm:ml-16 sm:mt-0 sm:flex-none">
              <NavLink
                to="/newProject"
                type="button"
                className="block rounded-md bg-gray-800 border border-gray-300 px-3 py-2 text-center text-sm font-semibold text-white hover:bg-gray-700 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500"
              >
                Add new project
              </NavLink>
            </div>
          </div>

          <div className="mt-8 flow-root">{/* <ProjectListContainer /> */}</div>
        </div>
      </div>
    </>
  );
}
