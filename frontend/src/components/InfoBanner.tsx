import { Banner } from "flowbite-react";
import { HiX } from "react-icons/hi";
import { MdInfo } from "react-icons/md";

interface InfoBannerProps {
  info: string;
  onAbort: () => void;
}
export default function InfoBanner({ info, onAbort }: InfoBannerProps) {
  return (
    <Banner>
      <div className="flex w-full justify-between rounded-md border border-teal-700 bg-gray-50 p-2 dark:border-teal-400/30 dark:bg-gray-700 ">
        <div className="ml-4 flex items-center">
          <p className="flex items-center text-sm font-normal text-gray-500 dark:text-gray-400">
            {/* <MdAnnouncement className="mr-4 h-4 w-4" /> */}
            <MdInfo className="mr-4 h-5 w-5 text-teal-500" />
            <span className="[&_p]:inline">{info}</span>
          </p>
        </div>
        <Banner.CollapseButton
          color="gray"
          className="border-0 bg-transparent text-gray-500 dark:text-gray-400"
          onAbort={onAbort}
        >
          <HiX className="h-4 w-4" />
        </Banner.CollapseButton>
      </div>
    </Banner>
  );
}
