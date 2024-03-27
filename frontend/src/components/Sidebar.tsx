import { NavLink } from "react-router-dom";
import classnames from "classnames";
export default function Sidebar() {
  return (
    <>
      <div className="w-full">
        <ul className="flex flex-col h-full w-full">
          <li className="w-full">
            <NavLink
              to="/projects"
              className={({ isActive }) =>
                classnames(
                  isActive
                    ? " border-l-4 border-teal-500  pl-2 text-white bg-gray-800"
                    : "pl-3 text-gray-300",
                  "inline-block py-2 w-full capitalized font-normal text-sm  hover:text-white"
                )
              }
            >
              Projects
            </NavLink>
          </li>
          <li className="w-full">
            <NavLink
              to="/sandbox"
              className={({ isActive }) =>
                classnames(
                  isActive
                    ? " border-l-4 border-teal-500  pl-2 text-white bg-gray-800"
                    : "pl-3 text-gray-300",
                  " inline-block w-full py-2 capitalized font-normal text-sm  hover:text-white"
                )
              }
            >
              Sandbox
            </NavLink>
            <li className="w-full">
              <NavLink
                to="/feedback"
                className={({ isActive }) =>
                  classnames(
                    isActive
                      ? " border-l-4 border-teal-500  pl-2 text-white bg-gray-800"
                      : "pl-3 text-gray-300",
                    " inline-block w-full py-2 capitalized font-normal text-sm  hover:text-white"
                  )
                }
              >
                Feedback
              </NavLink>
            </li>{" "}
          </li>
          <li className="w-full">
            <NavLink
              onClick={(e) => e.preventDefault()}
              to="/containers"
              className={({ isActive }) =>
                classnames(
                  isActive
                    ? " border-l-4 border-teal-500  pl-2 text-white bg-gray-800"
                    : "pl-3 text-gray-300",
                  // " inline-block w-full py-2 capitalized font-normal text-sm  hover:text-white"
                  " inline-block w-full py-2 capitalized font-normal text-sm text-gray-600 cursor-not-allowed"
                )
              }
            >
              Containers (coming soon)
            </NavLink>
          </li>
          <li className="w-full">
            <NavLink
              onClick={(e) => e.preventDefault()}
              to="/containers"
              className={({ isActive }) =>
                classnames(
                  isActive
                    ? " border-l-4 border-teal-500  pl-2 text-white bg-gray-800"
                    : "pl-3 text-gray-300",
                  // " inline-block w-full py-2 capitalized font-normal text-sm  hover:text-white"
                  " inline-block w-full py-2 capitalized font-normal text-sm text-gray-600 cursor-not-allowed"
                )
              }
            >
              Settings (coming soon)
            </NavLink>
          </li>
        </ul>
      </div>
    </>
  );
}
