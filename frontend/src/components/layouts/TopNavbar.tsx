import { Container, Navbar } from "react-bootstrap";
import { Link } from "react-router-dom";
import { NavbarWrapper } from "./styles";

const TopNavbar = () => {
  return (
    <NavbarWrapper>
      <Container fluid>
        <Navbar.Toggle />
        <Navbar.Collapse className="justify-content-end">
          <Link to="/" className="btn btn-primary">
            Log out
          </Link>
        </Navbar.Collapse>
      </Container>
    </NavbarWrapper>
  );
};

export default TopNavbar;
