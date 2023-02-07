import { useState } from "react";
import {
  Badge,
  Button,
  Col,
  Form,
  Image,
  InputGroup,
  Row,
} from "react-bootstrap";
import {
  KeyBox,
  KeyPairs,
  KeyPairsSection,
  StepBoxWrapper,
  StepHeader,
  StepBody,
  InputGroupWrapper,
} from "./styles";
import { PageHeader } from "./../../../components/layouts/styles";

// Images
import dsPadlockImg from "../../../assets/images/ds-padlock.png";
import eyeIcon from "../../../assets/images/eye-line.svg";
import eyeCloseIcon from "../../../assets/images/eye-close-line.svg";

export type Keys = {
  privateKey: string;
  publicKey: string;
};

const ManageKeyPairPage = () => {
  const [keys] = useState<Keys>();
  const [showPublicKey, setShowPublicKey] = useState(false);
  const [showPrivateKey, setShowPrivateKey] = useState(false);

  return (
    <>
      <PageHeader>
        <h2>
          <Image className="titleIcon" src={dsPadlockImg} />
          <span>Manage Key Pair</span>
        </h2>
      </PageHeader>
      <StepBoxWrapper className="active">
        <StepHeader>
          <h3 className="stepTitle">Current key pair</h3>
        </StepHeader>
        <StepBody>
          <Row>
            <Col md={12}>
              <InputGroupWrapper>
                <Form.Label>Private Key</Form.Label>
                <InputGroup className="mb-3">
                  <Form.Control
                    placeholder="Generated Private Key"
                    aria-label="Generated Private Key"
                    className={showPrivateKey ? "show" : "hide"}
                    defaultValue={keys?.privateKey}
                    readOnly={true}
                  />
                  <Button
                    variant="outline-secondary"
                    onClick={() => setShowPrivateKey(!showPrivateKey)}
                  >
                    <Image src={!showPrivateKey ? eyeIcon : eyeCloseIcon} />
                    {!showPrivateKey ? "SHOW" : "HIDE"}
                  </Button>
                </InputGroup>
              </InputGroupWrapper>
            </Col>
            <Col md={12}>
              <InputGroupWrapper>
                <Form.Label>Public Key</Form.Label>
                <InputGroup className="mb-3">
                  <Form.Control
                    placeholder="Generated Public Key"
                    aria-label="Generated Public Key"
                    className={showPublicKey ? "show" : "hide"}
                    defaultValue={keys?.publicKey}
                    readOnly={true}
                  />
                  <Button
                    variant="outline-secondary"
                    onClick={() => setShowPublicKey(!showPublicKey)}
                  >
                    <Image src={!showPublicKey ? eyeIcon : eyeCloseIcon} />
                    {!showPublicKey ? "SHOW" : "HIDE"}
                  </Button>
                </InputGroup>
              </InputGroupWrapper>
            </Col>
          </Row>
        </StepBody>
      </StepBoxWrapper>
      <StepBoxWrapper className="active">
        <StepHeader>
          <h3 className="stepTitle">Generate key pair</h3>
        </StepHeader>
        <StepBody>
          <Row>
            <Col md={12} lg={6}>
              <KeyPairsSection className="separator">
                <h4>Plain text keys</h4>
                <KeyPairs>
                  <KeyBox>
                    <span>RSA</span>
                    <input type="radio" name="radio" />
                    <span className="checkmark"></span>
                  </KeyBox>
                  <KeyBox>
                    <span>ECIES</span>
                    <input type="radio" name="radio" />
                    <span className="checkmark"></span>
                  </KeyBox>
                </KeyPairs>
              </KeyPairsSection>
            </Col>
            <Col md={12} lg={6}>
              <KeyPairsSection>
                <h4>
                  Non custodial keys <Badge bg="success">Premium</Badge>{" "}
                </h4>
                <KeyPairs>
                  <KeyBox>
                    <span>TKEY</span>
                    <input type="radio" name="radio" />
                    <span className="checkmark"></span>
                  </KeyBox>
                </KeyPairs>
              </KeyPairsSection>
            </Col>
          </Row>
        </StepBody>
      </StepBoxWrapper>
    </>
  );
};

export default ManageKeyPairPage;
