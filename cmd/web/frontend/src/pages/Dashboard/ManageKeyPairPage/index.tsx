import * as Yup from "yup";
import * as formik from "formik";
import { useEffect, useState } from "react";
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
  ColoredBtn,
} from "./styles";
import { PageHeader } from "./../../../components/layouts/styles";
import Select, { StylesConfig } from "react-select";

// Images
import dsPadlockImg from "../../../assets/images/ds-padlock.png";
import eyeIcon from "../../../assets/images/eye-line.svg";
import copyIcon from "../../../assets/images/file-copy-line.svg";
import eyeCloseIcon from "../../../assets/images/eye-close-line.svg";
import { GenerateKeyPair } from "../../../../wailsjs/go/main/App";
import {
  persistKekType,
  persistKey,
  readKekType,
  readKey,
} from "../../../services/localStorage.service";
import { types } from "../../../../wailsjs/go/models";
import { copyToClipboard } from "../../../helper";
import { toast } from "react-toastify";

const ManageKeyPairPage = () => {
  const [keys, setKeys] = useState<types.Keys>();
  const [currentKeys, setCurrentKeys] = useState<types.Keys>();
  const [kekType, setkekType] = useState<string>("");
  const [currentKekType, setCurrentKekType] = useState<string>("");
  const [showPublicKey, setShowPublicKey] = useState(false);
  const [showPrivateKey, setShowPrivateKey] = useState(false);
  const [isKeysGenerated, setIsKeysGenerated] = useState(false);
  const [showEditForm, setShowEditForm] = useState(false);
  const [tKeyMessage, setTKeyMessage] = useState<boolean>(false);

  const { Formik } = formik;

  const schema = Yup.object().shape({
    KekType: Yup.string().required("Please select kek type"),
    PublicKey: Yup.string().required("Please enter public key"),
    PrivateKey: Yup.string().required("Please enter private key"),
  });

  const generateKeyPair = async () => {
    const response = await GenerateKeyPair(kekType);
    if (response && response.Status == "success") {
      setKeys(response.Data);
      setIsKeysGenerated(!isKeysGenerated);
      persistKekType(kekType);

      toast.success("Key pair generated successfully.", {
        position: toast.POSITION.TOP_RIGHT,
      });
    }
  };

  useEffect(() => {
    setCurrentKeys(readKey());
    setCurrentKekType(readKekType());
  }, [setCurrentKeys, setCurrentKekType]);

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
          <span className="stepTitle">Current key pair</span>
          {!showEditForm && (
            <div className="right-part">
              <ColoredBtn
                className={`step-button ml-2`}
                onClick={() => setShowEditForm(!showEditForm)}
              >
                Edit
              </ColoredBtn>
            </div>
          )}
        </StepHeader>
        <StepBody>
          {showEditForm ? (
            <Formik
              validationSchema={schema}
              onSubmit={(data: any) => {
                persistKey({
                  PrivateKey: data.PrivateKey,
                  PublicKey: data.PublicKey,
                });
                persistKekType(data.KekType);
                setCurrentKeys({
                  PrivateKey: data.PrivateKey,
                  PublicKey: data.PublicKey,
                });
                setCurrentKekType(data.KekType);
                setShowEditForm(!showEditForm);
                toast.success("Key pair updated successfully.", {
                  position: toast.POSITION.TOP_RIGHT,
                });
              }}
              initialValues={{
                KekType: currentKekType,
                PublicKey: currentKeys?.PublicKey || "",
                PrivateKey: currentKeys?.PrivateKey || "",
              }}
            >
              {({
                handleSubmit,
                handleChange,
                handleBlur,
                values,
                touched,
                isValid,
                errors,
                setFieldValue,
              }) => (
                <Form noValidate onSubmit={handleSubmit}>
                  <Row>
                    <Col md={12} className="mb-3">
                      <Form.Group className="mb-3">
                        <Form.Label>Kek Type</Form.Label>
                        <Form.Select
                          name="KekType"
                          value={values.KekType}
                          onChange={handleChange}
                          isInvalid={!!errors.KekType}
                        >
                          <option>Please select kek type</option>
                          <option value="rsa">RSA</option>
                          <option value="ecies">ECIES</option>
                        </Form.Select>
                        <span
                          className="invalid-feedback"
                          style={{ color: "red", textAlign: "left" }}
                        >
                          {errors.KekType}
                        </span>
                      </Form.Group>
                    </Col>
                  </Row>
                  <Row>
                    <Col md={12} className="mb-3">
                      <Form.Label>Public Key</Form.Label>
                      <Form.Control
                        type="text"
                        placeholder="Public Key"
                        aria-label="Public Key"
                        name="PublicKey"
                        value={values.PublicKey}
                        onChange={handleChange}
                        isInvalid={!!errors.PublicKey}
                      />
                      <span
                        className="invalid-feedback"
                        style={{ color: "red", textAlign: "left" }}
                      >
                        {errors.PublicKey}
                      </span>
                    </Col>
                    <Col md={12}>
                      <Form.Label>Private Key</Form.Label>
                      <Form.Control
                        type="text"
                        placeholder="Private Key"
                        aria-label="Private Key"
                        name="PrivateKey"
                        value={values.PrivateKey}
                        onChange={handleChange}
                        isInvalid={!!errors.PrivateKey}
                      />
                      <span
                        className="invalid-feedback"
                        style={{ color: "red", textAlign: "left" }}
                      >
                        {errors.PrivateKey}
                      </span>
                    </Col>
                    <Col md={12}>
                      <ColoredBtn
                        type="button"
                        className="submitBtn"
                        onClick={() => {
                          setShowEditForm(!showEditForm);
                        }}
                        style={{ marginLeft: 5 }}
                      >
                        Cancel
                      </ColoredBtn>
                      <ColoredBtn type="submit" className="submitBtn">
                        Save Changes
                      </ColoredBtn>
                    </Col>
                  </Row>
                </Form>
              )}
            </Formik>
          ) : (
            <>
              <Row>
                <Col md={12}>
                  <Form.Label>KEK Type</Form.Label>
                  <InputGroupWrapper>
                    <InputGroup className="mb-3">
                      <Form.Control
                        placeholder="Enter your kek type"
                        defaultValue={currentKekType.toUpperCase()}
                        value={currentKekType.toUpperCase()}
                        disabled={true}
                      />
                      <Button
                        variant="outline-primary"
                        onClick={() => copyToClipboard(currentKekType || "")}
                      >
                        <Image src={copyIcon} />
                        Copy
                      </Button>
                    </InputGroup>
                  </InputGroupWrapper>
                </Col>
              </Row>
              <Row>
                <Col md={12}>
                  <Form.Label>Public Key</Form.Label>
                  <InputGroupWrapper>
                    <InputGroup className="mb-3">
                      <Form.Control
                        placeholder="Enter your public key"
                        defaultValue={currentKeys?.PublicKey}
                        value={currentKeys?.PublicKey}
                        disabled={true}
                      />
                      <Button
                        variant="outline-primary"
                        onClick={() =>
                          copyToClipboard(currentKeys?.PublicKey || "")
                        }
                      >
                        <Image src={copyIcon} />
                        Copy
                      </Button>
                    </InputGroup>
                  </InputGroupWrapper>
                </Col>
              </Row>
              <Row>
                <Col md={12}>
                  <Form.Label>Private Key</Form.Label>
                  <InputGroupWrapper>
                    <InputGroup className="mb-3">
                      <Form.Control
                        placeholder="Enter your private key"
                        defaultValue={currentKeys?.PrivateKey}
                        value={currentKeys?.PrivateKey}
                        disabled={true}
                      />
                      <Button
                        variant="outline-primary"
                        onClick={() =>
                          copyToClipboard(currentKeys?.PrivateKey || "")
                        }
                      >
                        <Image src={copyIcon} />
                        Copy
                      </Button>
                    </InputGroup>
                  </InputGroupWrapper>
                </Col>
              </Row>
            </>
          )}
        </StepBody>
      </StepBoxWrapper>
      <StepBoxWrapper className="active">
        <StepHeader>
          <span className="stepTitle">Generate key pair</span>
          {kekType !== "tkey" && (
            <div className="right-part">
              <ColoredBtn
                className={`step-button ml-2`}
                onClick={generateKeyPair}
              >
                Generate
              </ColoredBtn>
            </div>
          )}
        </StepHeader>
        <StepBody>
          <Row>
            <Col md={12} lg={6}>
              <KeyPairsSection className="separator">
                <h5>Plain text keys</h5>
                <KeyPairs>
                  <KeyBox>
                    <span>RSA</span>
                    <input
                      type="radio"
                      name="kekType"
                      checked={kekType === "rsa"}
                      onChange={() => {
                        setkekType("rsa");
                        setIsKeysGenerated(false);
                        setTKeyMessage(false);
                      }}
                    />
                    <span className="checkmark"></span>
                  </KeyBox>
                  <KeyBox>
                    <span>ECIES</span>
                    <input
                      type="radio"
                      name="kekType"
                      checked={kekType === "ecies"}
                      onChange={() => {
                        setkekType("ecies");
                        setIsKeysGenerated(false);
                        setTKeyMessage(false);
                      }}
                    />
                    <span className="checkmark"></span>
                  </KeyBox>
                </KeyPairs>
              </KeyPairsSection>
            </Col>
            <Col md={12} lg={6}>
              <KeyPairsSection>
                <h5>
                  Non custodial keys <Badge bg="success">Premium</Badge>{" "}
                </h5>
                <KeyPairs>
                  <KeyBox>
                    <span>TKEY</span>
                    <input
                      type="radio"
                      name="kekType"
                      checked={kekType === "tkey"}
                      onChange={() => {
                        setkekType("tkey");
                        setTKeyMessage(true);
                      }}
                    />
                    <span className="checkmark"></span>
                  </KeyBox>
                </KeyPairs>
                {tKeyMessage && (
                  <span>
                    Please purchase premium membership to use this feature.
                  </span>
                )}
              </KeyPairsSection>
            </Col>
          </Row>
        </StepBody>
      </StepBoxWrapper>
      {isKeysGenerated && (
        <StepBoxWrapper className="active">
          <StepHeader>
            <span className="stepTitle">Key pair generated!</span>
            <div className="right-part">
              <ColoredBtn
                className={`step-button ml-2`}
                onClick={() => {
                  const res = persistKey(new types.Keys(keys));
                  setCurrentKeys(res);
                  setCurrentKekType(kekType);
                  setIsKeysGenerated(false);
                  setkekType("");
                  setKeys(undefined);

                  toast.success("Key pair set successfully.", {
                    position: toast.POSITION.TOP_RIGHT,
                  });
                }}
              >
                Set as Current key pair
              </ColoredBtn>
            </div>
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
                      defaultValue={keys?.PrivateKey}
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
                      defaultValue={keys?.PublicKey}
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
      )}
    </>
  );
};

export default ManageKeyPairPage;
