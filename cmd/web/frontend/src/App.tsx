import "./App.css";
import { ThemeProvider } from "react-bootstrap";
import { Routes, Route } from "react-router-dom";
import LoginPage from "./pages/Auth/Login/LoginPage";
import Dashboard from "./pages/Dashboard";
import "bootstrap/dist/css/bootstrap.min.css";

function App() {
  return (
    <ThemeProvider
      breakpoints={["xxxl", "xxl", "xl", "lg", "md", "sm", "xs", "xxs"]}
      minBreakpoint="xxs"
    >
      <Routes>
        <Route path="/" element={<LoginPage />} />
        <Route path="/dashboard/*" element={<Dashboard />} />
      </Routes>
    </ThemeProvider>
  );
}

export default App;
