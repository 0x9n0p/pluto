import { useEffect } from "react";
import SalesProjectCard from "./SalesProjectCard";

export default function SalesProjectsList() {
  return (
    <div className="flex flex-col col-gap-2 min-h-full w-[400px]">
      <SalesProjectCard />
    </div>
  );
}
