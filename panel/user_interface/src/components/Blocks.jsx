import BlocksHeader from "./BlocksHeader";
import BlocksList from "./BlocksList";

function Blocks() {
  return (
    <div className="flex flex-col gap-8 w-1/2">
      <BlocksHeader />
      <BlocksList />
    </div>
  );
}

export default Blocks;
