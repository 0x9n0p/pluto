import { useEffect } from "react";
import SalesProjectCard from "./SalesProjectCard";

export default function SalesProjectsList() {
  let container = null;

  useEffect(() => {
    container = document.querySelector(".my-sales-project");
    console.log(container);
  }, []);

  function handleDragOver(event) {
    // Reset Action
    event.preventDefault();
  }

  function handleDrop(event) {
    event.preventDefault();

    // Select Box That Drag
    let targetId = event.dataTransfer.getData("boxId");
    let targetElem = document.getElementById(targetId);

    // Check if the dragged element exists
    // Create a new container div
    let saleBox = document.createElement("div");
    saleBox.classList.add("my-sales-project__content");

    // Create a divide span
    let divide = document.createElement("span");
    divide.classList.add("divide");

    // Append the dragged element and the divide span to the new container
    saleBox.append(targetElem);
    saleBox.append(divide);

    // Append the new container to the target container
    container.append(saleBox);
  }

  return (
    <div
      className="flex flex-col col-gap-2 min-h-full w-[400px] my-sales-project"
      onDragOver={handleDragOver}
      onDrop={handleDrop}
    >
      {/* Make sure to add draggable="true" to enable dragging */}
      <SalesProjectCard draggable="true" />
    </div>
  );
}
