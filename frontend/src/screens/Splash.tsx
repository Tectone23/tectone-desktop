import { useEffect, useState } from "react";
import Logo from "../assets/images/White_Logo.webp";
import { LoadConfig } from "../../wailsjs/go/main/Project";
import { CheckDependencies } from "../../wailsjs/go/main/App";
import { Navigate, createSearchParams, useNavigate } from "react-router-dom";
import { EventsOn } from "../../wailsjs/runtime/runtime";
// import { useApplictionStore } from "../store/store";
import { useProjects } from "../hooks/useProjects";

export default function Splash() {
  const navigate = useNavigate();
  const [deps, setDeps] = useState<string[]>([]);
  const [status, setStatus] = useState<string>("");

  // const getProjects = useApplictionStore((state) => state.getProjects);

  EventsOn("dependency-not-found", (dep) => {
    setDeps((state) => {
      const newState = [...state, dep];
      return newState;
    });
  });

  useEffect(() => {
    async function init() {
      setStatus("Loading configuration ...");
      await LoadConfig();
      setStatus("Checking for dependencies ...");
      await CheckDependencies();

      async function delay(ms: number) {
        await (async () => {
          return new Promise((resolve) => setTimeout(resolve, ms));
        })();
      }
      // getProjects();
      await delay(2000);

      // navigate to projects screen
      navigate({
        pathname: "projects",
        search: createSearchParams({
          missingDeps: deps,
        }).toString(),
      });
    }

    // init initialises the application
    init();
  }, []);

  return (
    <div className="flex w-full h-full items-center justify-center">
      <div className="flex flex-col space-y-8">
        <div className="flex flex-col items-center justify-center">
          <img src={Logo} alt="Tectone" />
          <span className="ml-24 uppercase -mt-4 font-bold text-gray-600">
            privacy on the move
          </span>
        </div>
        <div className="flex mx-auto text-gray-400 uppercase text-sm font-semibold">
          Loading ...
        </div>
      </div>
    </div>
  );
}
