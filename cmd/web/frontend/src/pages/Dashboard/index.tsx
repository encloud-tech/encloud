import { Route, Routes } from "react-router-dom";
import MainLayout from "../../components/layouts/MainLayout";
import GetResultsPage from "./Compute/GetResultsPage";
import ListComputePage from "./Compute/ListComputePage";
import ManageComputePage from "./Compute/ManageComputePage";
import ConfigurationPage from "./ConfigurationPage";
import ManageKeyPairPage from "./ManageKeyPairPage";
import ListContentPage from "./ListContentPage";
import UploadContent from "./UploadContentPage";
import RetrieveContentPage from "./RetrieveContentPage";
import RetrieveSharedContentPage from "./RetrieveSharedContentPage";

const Dashboard = () => {
  return (
    <MainLayout>
      <Routes>
        <Route path="/" element={<ManageKeyPairPage />} />
        <Route path="/upload" element={<UploadContent />} />
        <Route path="/list" element={<ListContentPage />} />
        <Route path="/retrieve/:id" element={<RetrieveContentPage />} />
        <Route
          path="/retrieve-shared-content"
          element={<RetrieveSharedContentPage />}
        />
        <Route path="/manage-compute" element={<ManageComputePage />} />
        <Route path="/list-compute" element={<ListComputePage />} />
        <Route path="/get-results/:id" element={<GetResultsPage />} />
        <Route path="/configuration" element={<ConfigurationPage />} />
      </Routes>
    </MainLayout>
  );
};

export default Dashboard;
