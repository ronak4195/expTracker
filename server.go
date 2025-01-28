import React, { useEffect, useState } from "react";

// Component to render UI based on the server response
const RenderComponent = ({ component }) => {
  switch (component.type) {
    case "text":
      return (
        <div style={{ fontSize: component.properties.style === "header" ? "24px" : "16px" }}>
          {component.properties.text}
        </div>
      );

    case "button":
      return (
        <button
          onClick={() => {
            fetch(`http://localhost:8080${component.properties.onClick}`)
              .then((response) => response.json())
              .then((data) => alert(data.message))
              .catch((error) => console.error("Error:", error));
          }}
          style={{
            padding: "10px 20px",
            backgroundColor: "#007BFF",
            color: "#FFF",
            border: "none",
            borderRadius: "5px",
            cursor: "pointer",
          }}
        >
          {component.properties.text}
        </button>
      );

    case "image":
      return (
        <img
          src={component.properties.src}
          alt={component.properties.alt}
          height={component.properties.height}
          style={{ margin: "20px 0" }}
        />
      );

    case "container":
      return (
        <div
          style={{
            display: "flex",
            flexDirection: component.properties.direction === "vertical" ? "column" : "row",
            padding: `${component.properties.padding}px`,
          }}
        >
          {component.children.map((child, index) => (
            <RenderComponent key={index} component={child} />
          ))}
        </div>
      );

    default:
      return <div>Unknown component type: {component.type}</div>;
  }
};

// Main App Component
const App = () => {
  const [uiData, setUiData] = useState(null);

  useEffect(() => {
    // Fetch UI data from the backend
    fetch("http://localhost:8080")
      .then((response) => response.json())
      .then((data) => setUiData(data))
      .catch((error) => console.error("Error fetching UI data:", error));
  }, []);

  return (
    <div style={{ fontFamily: "Arial, sans-serif", margin: "20px" }}>
      {uiData ? (
        <RenderComponent component={uiData} />
      ) : (
        <div>Loading UI...</div>
      )}
    </div>
  );
};

export default App;
