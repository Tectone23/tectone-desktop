import { useCallback, useState } from "react";
import { GetAllExampleProjectNames } from "../../wailsjs/go/main/SDK";
import { model } from "../../wailsjs/go/models";

export function useSdk() {
  const [projectNames, setProjectNames] = useState<string[] | null>(null);

  const getExampleProjectNames = useCallback(
    async (path: string) => {
      const proj = await GetAllExampleProjectNames(path);
      setProjectNames(proj);
    },
    [setProjectNames]
  );

  return { projectNames, getExampleProjectNames };
}
