import Sidebar from "../components/Sidebar";
import Logo from "../assets/images/White_Logo.webp";
import ErrorBanner from "../components/ErrorBanner";
import { useError, useInfo } from "../hooks/useMessaging";
import InfoBanner from "../components/InfoBanner";

interface ContainerProps {
  children: React.ReactNode | React.ReactNode[];
}
export default function Container({ children }: ContainerProps) {
  return (
    <>
      <div className="w-[300px] max-h-full  border-r border-gray-800 pt-12  bg-gray-900 h-full px-6">
        <div className="flex justify-between items-center  pb-12 border-b border-gray-700 mb-12">
          <img src={Logo} alt="" />
        </div>
        <Sidebar />
      </div>
      <div className="flex-col pt-20 space-y-8 relative flex w-full max-h-full overflow-y-scroll text-white px-8">
        {children}
      </div>
    </>
  );
}
