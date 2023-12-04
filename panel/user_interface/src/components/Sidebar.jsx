import Logo from "./Logo";
import LogoutBtn from "./LogoutBtn";
import Menu from "./Menu";

function Sidebar() {
  return (
    <aside className="py-4 px-0 bg-white flex flex-col w-[200px] gap-6 relative border-r border-gray-300">
      <Logo />
      <Menu />
      <LogoutBtn />
    </aside>
  );
}

export default Sidebar;
