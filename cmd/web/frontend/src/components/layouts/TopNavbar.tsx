import { Container, Navbar } from "react-bootstrap";
import { Link } from "react-router-dom";
import { NavbarWrapper } from "./styles";

const TopNavbar = () => {
  return (
    <NavbarWrapper>
      <Container fluid>
        <Navbar.Toggle />
      </Container>
    </NavbarWrapper>
  );
};

export default TopNavbar;
