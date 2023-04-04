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
  OverlayTrigger,
  Row,
  Tooltip,
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
import Select, { StylesConfig } from "react-select";

// Images
import padlockIcon from "../../../assets/images/padlock.png";
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

const colourStyles: StylesConfig = {
  control: (styles, state) => ({
    ...styles,
    backgroundColor: "white",
    borderColor: state.isFocused ? "#cc336610" : "#cc336610",
    boxShadow: "0 0 0 0px #cc336610",
    ":hover": {
      ...styles[":hover"],
      borderColor: "#cc336610",
      boxShadow: "0 0 0 0px #cc336610",
    },
    ":focus": {
      ...styles[":focus"],
      borderColor: "#cc336610",
      boxShadow: "0 0 0 0px #cc336610",
    },
    ":active": {
      ...styles[":active"],
      borderColor: "#cc336610",
      boxShadow: "0 0 0 0px #cc336610",
    },
  }),
  option: (styles, { data, isDisabled, isFocused, isSelected }) => {
    return {
      ...styles,
      // backgroundColor: "pink",
      color: "black",
      backgroundColor: isSelected ? "#d94964" : "#ffffff",

      ":active": {
        ...styles[":active"],
        backgroundColor: !isDisabled
          ? isSelected
            ? "#d94964"
            : "red"
          : undefined,
      },
      ":hover": {
        ...styles[":hover"],
        backgroundColor: isSelected ? "#d94964" : "#d949643b",
      },
    };
  },
  input: (styles) => ({ ...styles, borderRadius: ".375rem" }),
  placeholder: (styles) => ({ ...styles }),
  singleValue: (styles, { data }) => ({ ...styles }),
};

const kekTypeOptions = [
  { value: "rsa", label: "RSA" },
  { value: "ecies", label: "ECIES" },
];

const renderTooltip = (props: any) => (
  <Tooltip id="button-tooltip" {...props}>
    <p>
      <strong>RSA</strong> - RSAES-OAEP 3072 bit key with a SHA-256 digest
    </p>
    <p>
      <strong>ECIES</strong> - The ECIES standard combines ECC-based asymmetric
      cryptography with symmetric ciphers. ECC is the modern and the preferable
      public-key cryptosystem due to smaller keys, shorter signatures and better
      performance.
    </p>
  </Tooltip>
);

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
  const [kekErrorMessage, setKekErrorMessage] = useState<boolean>(false);

  const { Formik } = formik;

  const schema = Yup.object().shape({
    PublicKey: Yup.string().required("Please enter public key"),
    PrivateKey: Yup.string().required("Please enter private key"),
  });

  const generateKeyPair = async () => {
    if (kekType) {
      setKekErrorMessage(false);
      const response = await GenerateKeyPair(kekType);
      if (response && response.Status == "success") {
        setKeys(response.Data);
        setIsKeysGenerated(!isKeysGenerated);
        persistKekType(kekType);

        toast.success("Key pair generated successfully.", {
          position: toast.POSITION.TOP_RIGHT,
        });
      } else {
        toast.error("Something went wrong!.Please retry", {
          position: toast.POSITION.TOP_RIGHT,
        });
      }
    } else {
      setKekErrorMessage(true);
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
          <Image className="titleIcon" src={padlockIcon} />
          <span>Manage Key Pair</span>
        </h2>
      </PageHeader>
      <StepBoxWrapper className="active">
        <StepHeader>
          <span className="stepTitle">Current key pair</span>
          {!showEditForm && (
            <div className="right-part">
              <Button onClick={() => setShowEditForm(!showEditForm)}>
                Edit
              </Button>
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
                persistKekType(data.KekType.value);
                setCurrentKeys({
                  PrivateKey: data.PrivateKey,
                  PublicKey: data.PublicKey,
                });
                setCurrentKekType(data.KekType.value);
                setShowEditForm(!showEditForm);
                toast.success("Key pair updated successfully.", {
                  position: toast.POSITION.TOP_RIGHT,
                });
              }}
              initialValues={{
                KekType:
                  currentKekType === "rsa"
                    ? kekTypeOptions[0]
                    : kekTypeOptions[1],
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
                        <Form.Label>
                          Kek Type
                          <OverlayTrigger
                            placement="right"
                            delay={{ show: 250, hide: 400 }}
                            overlay={renderTooltip}
                          >
                            <i
                              style={{ marginLeft: 3 }}
                              className="fa fa-info-circle"
                              aria-hidden="true"
                            ></i>
                          </OverlayTrigger>
                        </Form.Label>
                        <Select
                          name="KekType"
                          className="dek-type-select"
                          styles={colourStyles}
                          options={kekTypeOptions}
                          value={values.KekType}
                          onChange={(newVal) => {
                            setFieldValue("KekType", newVal);
                          }}
                        />
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
                        className="editInput"
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
                        className="editInput"
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
                    <Col md={12} className="buttoncolumn">
                      <Button
                        type="button"
                        onClick={() => {
                          setShowEditForm(!showEditForm);
                        }}
                        style={{ marginRight: 15 }}
                      >
                        Cancel
                      </Button>
                      <Button type="submit">Save Changes</Button>
                    </Col>
                  </Row>
                </Form>
              )}
            </Formik>
          ) : (
            <>
              <Row>
                <Col md={12}>
                  <Form.Label>
                    KEK Type
                    <OverlayTrigger
                      placement="right"
                      delay={{ show: 250, hide: 400 }}
                      overlay={renderTooltip}
                    >
                      <i
                        style={{ marginLeft: 3 }}
                        className="fa fa-info-circle"
                        aria-hidden="true"
                      ></i>
                    </OverlayTrigger>
                  </Form.Label>
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
              <Button onClick={generateKeyPair}>Generate</Button>
            </div>
          )}
        </StepHeader>
        <StepBody>
          <Row>
            <Col md={12} lg={6}>
              <KeyPairsSection className="separator">
                <h5>
                  Plain text keys
                  <OverlayTrigger
                    placement="right"
                    delay={{ show: 250, hide: 400 }}
                    overlay={renderTooltip}
                  >
                    <i
                      style={{ marginLeft: 5 }}
                      className="fa fa-info-circle"
                      aria-hidden="true"
                    ></i>
                  </OverlayTrigger>
                </h5>
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
                        setKekErrorMessage(false);
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
                        setKekErrorMessage(false);
                      }}
                    />
                    <span className="checkmark"></span>
                  </KeyBox>
                </KeyPairs>
                {kekErrorMessage && (
                  <span>Please select kek type to generate key</span>
                )}
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
                        setKekErrorMessage(false);
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
              <Button
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
              </Button>
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
