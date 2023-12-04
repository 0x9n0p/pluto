import AddPiplineBtn from "./AddPiplineBtn";

export default function Header() {
  return (
    <header className="flex flex-col gap-8 items-baseline">
      <h1 className="text-3xl font-medium">Pipline Name</h1>
      {/* <!-- Button For Add Pipline --> */}
      <AddPiplineBtn />
    </header>
  );
}
