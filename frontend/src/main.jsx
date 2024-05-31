import App from "./App";
import { BrowserRouter as Router } from "react-router-dom";
import { createRoot } from "react-dom/client";
import "./index.css";

const domNode = document.getElementById("root");
const root = createRoot(domNode);

root.render(
  <Router>
    <App />
  </Router>
);
