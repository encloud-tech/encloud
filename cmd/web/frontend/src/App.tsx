import "./App.css";
import { ThemeProvider } from "react-bootstrap";
import { Routes, Route } from "react-router-dom";
import Dashboard from "./pages/Dashboard";
import "bootstrap/dist/css/bootstrap.min.css";

function App() {
  return (
    <ThemeProvider
      breakpoints={["xxxl", "xxl", "xl", "lg", "md", "sm", "xs", "xxs"]}
      minBreakpoint="xxs"
    >
      <Routes>
        <Route path="/*" element={<Dashboard />} />
      </Routes>
    </ThemeProvider>
  );
}

export default App;
