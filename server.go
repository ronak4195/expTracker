import React, { useEffect, useState } from "react";
import axios from "axios";

// Component to render UI components dynamically
const UIComponentRenderer = ({ component }) => {
  const { type, properties, children } = component;

  switch (type) {
    case "container":
      return (
        <div
          style={{
            display: "flex",
            flexDirection: properties.direction || "column",
            padding: properties.padding || 0,
          }}
        >
          {children.map((child, index) => (
            <UIComponentRenderer key={index} component={child} />
          ))}
        </div>
      );

    case "text":
      return (
        <h1 style={properties.style === "header" ? { fontSize: "24px", fontWeight: "bold" } : {}}>
          {properties.text}
        </h1>
      );

    case "button":
      return (
        <button
          onClick={() => {
            if (properties.onClick) {
              axios.get(properties.onClick).then((response) => {
                alert(response.data.message);
              });
            }
          }}
        >
          {properties.text}
        </button>
      );

    case "image":
      return (
        <img
          src={properties.src}
          alt={properties.alt}
          style={{ height: properties.height || "auto" }}
        />
      );

    default:
      return null;
  }
};

// Main App Component
const App = () => {
  const [uiData, setUiData] = useState(null);

  // Fetch UI data from the Golang server
  useEffect(() => {
    axios
      .get("http://localhost:8080/")
      .then((response) => {
        setUiData(response.data);
      })
      .catch((error) => {
        console.error("Error fetching UI data:", error);
      });
  }, []);

  if (!uiData) {
    return <div>Loading...</div>;
  }

  return (
    <div className="App">
      <UIComponentRenderer component={uiData} />
    </div>
  );
};

export default App;
