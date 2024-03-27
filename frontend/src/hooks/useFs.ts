import { useCallback, useEffect, useState } from "react";
import { OpenDir, OpenFile } from "../../wailsjs/go/main/App";

export function useDirDialog(): {
  projectDir: string;
  getProjectDir: () => void;
  filepath: string;
  getFilepath: () => void;
} {
  const [projectDir, setProjectDir] = useState<string>("");
  const [filepath, setFilepath] = useState<string>("");

  const getProjectDir = useCallback(async () => {
    const pd = await OpenDir();
    setProjectDir(pd);
  }, [setProjectDir]);

  const getFilepath = useCallback(async () => {
    const p = await OpenFile();
    setFilepath(p);
    console.log(`got file: ${p}`);
  }, [setFilepath]);

  return {
    projectDir,
    getProjectDir,
    filepath,
    getFilepath,
  };
}
