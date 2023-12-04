import BlocksSearchBar from "./BlocksSearchBar";

export default function BlocksHeader() {
  return (
    <div className="flex flex-col gap-6">
      <h1 className="text-2xl font-bold">Blocks</h1>
      <BlocksSearchBar />
    </div>
  );
}
