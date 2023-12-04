// import Header from "./Header";
import SalesProjectsList from "./SalesProjectsList";
import Blocks from "./Blocks";
import Header from "./Header";

function MainContent() {
  return (
    <main className="w-[100%] bg-[#F5F5F5] flex flex-col p-8 gap-10 min-h-screen">
      <Header />
      <div className="flex items-start gap-8 justify-center">
        {/* <SalesProjectsList /> */}
        <SalesProjectsList />
        <Blocks />
      </div>
    </main>
  );
}

export default MainContent;
