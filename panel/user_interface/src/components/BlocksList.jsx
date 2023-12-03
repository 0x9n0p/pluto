import { useEffect, useState } from "react";
import BlockBox from "./BlockBox";

export default function BlocksList() {
  const [blocks, setBlocks] = useState([]);

  useEffect(() => {
    const fetchBlocks = async () => {
      const response = await fetch(
        "http://panel.plutoengine.ir/api/v1/processors"
      );
      const data = await response.json();

      setBlocks(data);
    };
    fetchBlocks();
  }, [blocks]);

  return (
    <div
      className="flex flex-col gap-8
    "
    >
      {blocks.map((item, i) => {
        return (
          <BlockBox
            key={i}
            icon={item.icon}
            description={item.description}
            title={item.name}
          />
        );
      })}
    </div>
  );
}
