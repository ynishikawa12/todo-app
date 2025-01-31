import { StrictMode } from "react";
import ReactDom from "react-dom/client";
import App from "./routes/App.tsx";

ReactDom.createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <App />
  </StrictMode>,
);
