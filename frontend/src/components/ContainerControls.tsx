import { PlayIcon, StopIcon } from "@heroicons/react/20/solid";

interface ContainerControlsProps {
  onStart: () => void;
  onStop: () => void;
  status: string;
}
export default function ContainerControls({
  onStart,
  onStop,
  status,
}: ContainerControlsProps) {
  return (
    <>
      <div className="flex space-x-2">
        <div onClick={() => onStart()}>
          <PlayIcon
            className={`${
              status == "exited"
                ? "text-teal-500 hover:text-teal-300"
                : status == "starting" || status == "stopping"
                ? "text-orange-500 hover:text-orange-300"
                : "text-gray-600"
            } h-4 w-4`}
          />
        </div>
        <div onClick={() => onStop()}>
          <StopIcon
            className={`${
              status == "running"
                ? "text-teal-500 hover:text-teal-300"
                : status == "stopping" || status == "starting"
                ? "text-orange-500 hover:text-orange-300"
                : "text-gray-600"
            } h-4 w-4`}
          />
        </div>
      </div>
    </>
  );
}
