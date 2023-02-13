import { Image } from "react-bootstrap";
import { AuthWrapper } from "./LoginPage.styles";
import encloudLogo from "../../../assets/images/encloud-logo.png";
import { Link } from "react-router-dom";

const LoginPage = () => {
  return (
    <AuthWrapper>
      <div className="auth-inner">
        <form>
          <Image src={encloudLogo} className="siteLogo" fluid />
          <div className="mb-3">
            <label>Email address</label>
            <input
              type="email"
              className="form-control"
              placeholder="Enter email"
              defaultValue="hello@encloud.tech"
            />
          </div>
          <div className="mb-3">
            <label>Password</label>
            <input
              type="password"
              className="form-control"
              placeholder="Enter password"
              defaultValue="encloud@123"
            />
          </div>
          <div className="d-grid">
            <Link to="/dashboard">
              <button type="button" className="btn btn-primary">
                Submit
              </button>
            </Link>
          </div>
        </form>
      </div>
    </AuthWrapper>
  );
};

export default LoginPage;
