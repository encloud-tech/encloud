import { Route, Routes } from "react-router-dom";
import MainLayout from "../../components/layouts/MainLayout";
import GetResultsPage from "./Compute/GetResultsPage";
import ListComputePage from "./Compute/ListComputePage";
import ManageComputePage from "./Compute/ManageComputePage";
import ConfigurationPage from "./ConfigurationPage";
import ManageKeyPairPage from "./ManageKeyPairPage";
import RetrievePage from "./RetrievePage";
import UploadContent from "./UploadContentPage";

const Dashboard = () => {
  return (
    <MainLayout>
      <Routes>
        <Route path="/" element={<ManageKeyPairPage />} />
        <Route path="/upload" element={<UploadContent />} />
        <Route path="/retrieve" element={<RetrievePage />} />
        <Route path="/manage-compute" element={<ManageComputePage />} />
        <Route path="/list" element={<ListComputePage />} />
        <Route path="/get-results/:id" element={<GetResultsPage />} />
        <Route path="/configuration" element={<ConfigurationPage />} />
      </Routes>
    </MainLayout>
  );
};

export default Dashboard;
