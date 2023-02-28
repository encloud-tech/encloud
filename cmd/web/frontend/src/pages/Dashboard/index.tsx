import { Route, Routes } from "react-router-dom";
import MainLayout from "../../components/layouts/MainLayout";
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
        <Route path="/configuration" element={<ConfigurationPage />} />
      </Routes>
    </MainLayout>
  );
};

export default Dashboard;
