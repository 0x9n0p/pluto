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
      console.log(response);
    };
    fetchBlocks();
  }, [blocks]);

  return (
    <div
      className="flex flex-col gap-8
    "
    >
      {blocks &&
        blocks.map((item, i) => {
          return (
            <BlockBox
              key={i}
              icon={item.icon}
              description={item.description}
              title={item.name}
              id={`blocks__box--${i}`}
            />
          );
        })}
      <BlockBox
        icon={null}
        description="test description"
        title="test"
        id="s22"
      />
      <BlockBox
        icon={null}
        description="test description"
        title="test"
        id="f11"
      />
    </div>
  );
}
