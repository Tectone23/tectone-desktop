import { Banner } from "flowbite-react";
import { HiX } from "react-icons/hi";
import { MdAnnouncement } from "react-icons/md";
import { MdError } from "react-icons/md";

interface ErrorBannerProps {
  error: string;
  onAbort: () => void;
}

export default function ErrorBanner({ error, onAbort }: ErrorBannerProps) {
  return (
    <Banner>
      <div className="flex w-full justify-between rounded-md border-b border-gray-200 bg-gray-50 p-2 dark:border-gray-600 dark:bg-gray-700">
        <div className="mx-auto flex items-center">
          <p className="flex items-center text-sm font-normal text-gray-500 dark:text-gray-400">
            {/* <MdAnnouncement className="mr-4 h-4 w-4" /> */}
            <MdError className="mr-4 h-5 w-5 text-red-500" />
            <span className="[&_p]:inline">{error}</span>
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
