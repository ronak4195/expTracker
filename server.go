import React, { useEffect, useState } from "react";

const fetchUI = async () => {
  try {
    const response = await fetch("http://localhost:8080/");
    return response.json();
  } catch (error) {
    console.error("Error fetching UI:", error);
    return null;
  }
};

const handleClick = async (url) => {
  try {
    const response = await fetch(url);
    const data = await response.json();
    alert(data.message); // Show response message
  } catch (error) {
    console.error("Error handling click:", error);
  }
};

const renderComponent = (component) => {
  switch (component.type) {
    case "container":
      return (
        <div
          style={{
            display: "flex",
            flexDirection: component.properties.direction || "column",
            padding: component.properties.padding || 0,
          }}
        >
          {component.children &&
            component.children.map((child, index) => (
              <React.Fragment key={index}>
                {renderComponent(child)}
              </React.Fragment>
            ))}
        </div>
      );

    case "text":
      return (
        <h1 style={{ fontSize: component.properties.style === "header" ? "24px" : "16px" }}>
          {component.properties.text}
        </h1>
      );

    case "button":
      return (
        <button onClick={() => handleClick(component.properties.onClick)}>
          {component.properties.text}
        </button>
      );

    case "image":
      return (
        <img
          src={component.properties.src}
          alt={component.properties.alt}
          height={component.properties.height}
        />
      );

    default:
      return null;
  }
};

const App = () => {
  const [uiData, setUiData] = useState(null);

  useEffect(() => {
    fetchUI().then((data) => setUiData(data));
  }, []);

  if (!uiData) return <p>Loading...</p>;

  return <div>{renderComponent(uiData)}</div>;
};

export default App;
