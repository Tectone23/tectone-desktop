import Sidebar from "../components/Sidebar";
import Logo from "../assets/images/White_Logo.webp";

export default function Containers() {
  return (
    <>
      <div className="w-[300px] border-r border-gray-800 pt-8  bg-gray-900 h-full px-6">
        <div className="flex justify-between items-center  pb-12 border-b border-gray-700 mb-12">
          <img src={Logo} alt="" />
        </div>
        <Sidebar />
      </div>
    </>
  );
}
