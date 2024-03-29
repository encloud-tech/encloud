import * as formik from "formik";
import {
  Button,
  Card,
  Col,
  Form,
  Row,
  Image,
  Tooltip,
  OverlayTrigger,
} from "react-bootstrap";
import { ColoredBtn, KeyBoxedContent, KeyPairsSection } from "./styles";
import { PageHeader } from "./../../../components/layouts/styles";
import Select, { StylesConfig } from "react-select";
import { CSSProperties, useEffect, useState } from "react";
import {
  FetchConfig,
  RestoreDefaultConfig,
  StoreConfig,
} from "../../../../wailsjs/go/main/App";
import { toast } from "react-toastify";
import settingsIcon from "../../../assets/images/settings.png";
import ClipLoader from "react-spinners/ClipLoader";

const override: CSSProperties = {
  display: "block",
  margin: "0 auto",
  borderColor: "white",
};

const storageTypeOptions = [
  { value: "badgerdb", label: "Badger DB" },
  { value: "couchbase", label: "Couchbase DB" },
];

const kekTypeOptions = [
  { value: "rsa", label: "RSA" },
  { value: "ecies", label: "ECIES" },
];

const emailTypeOptions = [
  { value: "smtp", label: "SMTP" },
  { value: "mailersend", label: "Mailersend" },
];

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
  input: (styles) => ({ ...styles }),
  placeholder: (styles) => ({ ...styles }),
  singleValue: (styles, { data }) => ({ ...styles }),
};

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

const ConfigurationPage = () => {
  const { Formik } = formik;
  const [loading, setLoading] = useState(false);
  const [restoreDefaultLoading, setRestoreDefaultLoading] = useState(false);
  const [config, setConfigData] = useState({
    Estuary: {
      BaseApiUrl: "",
      UploadApiUrl: "",
      GatewayApiUrl: "",
      CdnApiUrl: "",
      Token: "",
    },
    Email: {
      From: "",
      EmailType: emailTypeOptions[0],
      SMTP: {
        Server: "",
        Port: "",
        Username: "",
        Password: "",
      },
    },
    Stat: {
      KekType: kekTypeOptions[0],
      StorageType: storageTypeOptions[0],
      BadgerDB: {
        Path: "",
      },
      Couchbase: {
        Host: "",
        Username: "",
        Password: "",
        Bucket: {
          Name: "",
          Scope: "",
          Collection: "",
        },
      },
    },
  });

  const fetchData = async () => {
    const response = await FetchConfig();

    if (response.Data) {
      response.Data.Stat.KekType =
        response.Data.Stat.KekType === "rsa"
          ? kekTypeOptions[0]
          : kekTypeOptions[1];
      response.Data.Stat.StorageType =
        response.Data.Stat.StorageType === "badgerdb"
          ? storageTypeOptions[0]
          : storageTypeOptions[1];
      response.Data.Email.EmailType =
        response.Data.Email.EmailType === "smtp"
          ? emailTypeOptions[0]
          : emailTypeOptions[1];
      setConfigData(response.Data);
    }
  };

  useEffect(() => {
    fetchData();
  }, [setConfigData]);

  const restoreDefault = () => {
    setRestoreDefaultLoading(true);
    try {
      RestoreDefaultConfig().then((response) => {
        if (response && response.Status == "success") {
          setRestoreDefaultLoading(false);
          toast.success(response.Message, {
            position: toast.POSITION.TOP_RIGHT,
          });

          fetchData();
        } else {
          setRestoreDefaultLoading(false);
          toast.error("Something went wrong!", {
            position: toast.POSITION.TOP_RIGHT,
          });
        }
      });
    } catch (error) {
      setRestoreDefaultLoading(false);
      toast.error("Something went wrong!", {
        position: toast.POSITION.TOP_RIGHT,
      });
    }
  };

  const save = async (data: any) => {
    setLoading(true);
    data.Stat.KekType = data.Stat.KekType.value;
    data.Stat.StorageType = data.Stat.StorageType.value;
    data.Email.EmailType = data.Email.EmailType.value;
    try {
      StoreConfig(data).then((response) => {
        if (response && response.Status == "success") {
          setLoading(false);
          toast.success(response.Message, {
            position: toast.POSITION.TOP_RIGHT,
          });

          fetchData();
        } else {
          setLoading(false);
          toast.error("Something went wrong!", {
            position: toast.POSITION.TOP_RIGHT,
          });
        }
      });
    } catch (error) {
      setLoading(false);
      toast.error("Something went wrong!", {
        position: toast.POSITION.TOP_RIGHT,
      });
    }
  };

  return (
    <>
      <PageHeader>
        <h2>
          <Image className="titleIcon" src={settingsIcon} />
          <span>Configuration</span>
        </h2>
      </PageHeader>
      <Card className="mt-4">
        <Card.Body>
          <KeyBoxedContent>
            <Formik
              onSubmit={save}
              initialValues={config}
              enableReinitialize={true}
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
                    <Col md={12} lg={6}>
                      <KeyPairsSection>
                        <h4>Estuary</h4>
                      </KeyPairsSection>
                    </Col>
                  </Row>
                  <Row className="mt-2">
                    <Col md={6}>
                      <Form.Group className="mb-3">
                        <Form.Label>API URL</Form.Label>
                        <Form.Control
                          type="text"
                          name="Estuary.BaseApiUrl"
                          placeholder="https://api.estuary.tech"
                          value={values.Estuary.BaseApiUrl}
                          onChange={handleChange}
                        />
                      </Form.Group>
                    </Col>
                    <Col md={6}>
                      <Form.Group className="mb-3">
                        <Form.Label>Upload API URL</Form.Label>
                        <Form.Control
                          type="text"
                          name="Estuary.UploadApiUrl"
                          placeholder="https://edge.estuary.tech/api/v1"
                          value={values.Estuary.UploadApiUrl}
                          onChange={handleChange}
                        />
                      </Form.Group>
                    </Col>
                  </Row>
                  <Row>
                    <Col md={6}>
                      <Form.Group className="mb-3">
                        <Form.Label>Gateway API URL</Form.Label>
                        <Form.Control
                          type="text"
                          name="Estuary.GatewayApiUrl"
                          placeholder="https://edge.estuary.tech"
                          value={values.Estuary.GatewayApiUrl}
                          onChange={handleChange}
                        />
                      </Form.Group>
                    </Col>
                    <Col md={6}>
                      <Form.Group className="mb-3">
                        <Form.Label>CDN API URL</Form.Label>
                        <Form.Control
                          type="text"
                          name="Estuary.CdnApiUrl"
                          placeholder="https://cdn.estuary.tech"
                          value={values.Estuary.CdnApiUrl}
                          onChange={handleChange}
                        />
                      </Form.Group>
                    </Col>
                  </Row>
                  <Row>
                    <Col md={6}>
                      <Form.Group className="mb-3">
                        <Form.Label>Token</Form.Label>
                        <Form.Control
                          type="text"
                          name="Estuary.Token"
                          placeholder="EST6315eb22-XXXX-XXXX-XXXX-1acb4a954070ARY"
                          value={values.Estuary.Token}
                          onChange={handleChange}
                        />
                      </Form.Group>
                    </Col>
                  </Row>
                  <Row>
                    <Col md={12} lg={6}>
                      <KeyPairsSection>
                        <h4>Email</h4>
                      </KeyPairsSection>
                    </Col>
                  </Row>
                  <Row className="mt-2">
                    <Col md={6}>
                      <Form.Group className="mb-3">
                        <Form.Label>Email Type</Form.Label>
                        <Select
                          className="dek-type-select"
                          styles={colourStyles}
                          options={emailTypeOptions}
                          name="Email.EmailType"
                          value={values.Email.EmailType}
                          onChange={(newVal) => {
                            setFieldValue("Email.EmailType", newVal);
                          }}
                        />
                      </Form.Group>
                    </Col>
                    <Col md={6}>
                      <Form.Group className="mb-3">
                        <Form.Label>From Address</Form.Label>
                        <Form.Control
                          type="text"
                          placeholder="contact@encloud.tech"
                          name="Email.From"
                          value={values.Email.From}
                          onChange={handleChange}
                        />
                      </Form.Group>
                    </Col>
                  </Row>
                  {values.Email.EmailType.value === "smtp" && (
                    <>
                      <Row>
                        <Col md={6}>
                          <Form.Group className="mb-3">
                            <Form.Label>Server</Form.Label>
                            <Form.Control
                              type="text"
                              placeholder="smtp.mailtrap.io"
                              name="Email.SMTP.Server"
                              value={values.Email.SMTP.Server}
                              onChange={handleChange}
                            />
                          </Form.Group>
                        </Col>
                        <Col md={6}>
                          <Form.Group className="mb-3">
                            <Form.Label>Port</Form.Label>
                            <Form.Control
                              type="text"
                              placeholder="2525"
                              name="Email.SMTP.Port"
                              value={values.Email.SMTP.Port}
                              onChange={handleChange}
                            />
                          </Form.Group>
                        </Col>
                      </Row>
                      <Row>
                        <Col md={6}>
                          <Form.Group className="mb-3">
                            <Form.Label>Username</Form.Label>
                            <Form.Control
                              type="text"
                              placeholder="ac984e52bfd98h"
                              name="Email.SMTP.Username"
                              value={values.Email.SMTP.Username}
                              onChange={handleChange}
                            />
                          </Form.Group>
                        </Col>
                        <Col md={6}>
                          <Form.Group className="mb-3">
                            <Form.Label>Password</Form.Label>
                            <Form.Control
                              type="password"
                              placeholder="861b495c076987"
                              name="Email.SMTP.Password"
                              value={values.Email.SMTP.Password}
                              onChange={handleChange}
                            />
                          </Form.Group>
                        </Col>
                      </Row>
                    </>
                  )}
                  <Row>
                    <Col md={12} lg={6}>
                      <KeyPairsSection>
                        <h4>Storage & KEK Type</h4>
                      </KeyPairsSection>
                    </Col>
                  </Row>
                  <Row className="mt-2">
                    <Col md={6}>
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
                          className="dek-type-select"
                          styles={colourStyles}
                          options={kekTypeOptions}
                          name="Stat.KekType"
                          value={values.Stat.KekType}
                          onChange={(newVal) => {
                            setFieldValue("Stat.KekType", newVal);
                          }}
                        />
                      </Form.Group>
                    </Col>
                    <Col md={6}>
                      <Form.Group className="mb-3">
                        <Form.Label>Storage Type</Form.Label>
                        <Select
                          className="dek-type-select"
                          styles={colourStyles}
                          options={storageTypeOptions}
                          name="Stat.StorageType"
                          value={values.Stat.StorageType}
                          onChange={(newVal) => {
                            setFieldValue("Stat.StorageType", newVal);
                          }}
                        />
                      </Form.Group>
                    </Col>
                  </Row>
                  {values.Stat.StorageType.value === "badgerdb" ? (
                    <Row className="mt-2">
                      <Col md={6}>
                        <Form.Group className="mb-3">
                          <Form.Label>Path</Form.Label>
                          <Form.Control
                            type="text"
                            placeholder="badger.db"
                            name="Stat.BadgerDB.Path"
                            value={values.Stat.BadgerDB.Path}
                            onChange={handleChange}
                          />
                        </Form.Group>
                      </Col>
                      <Col md={6}></Col>
                    </Row>
                  ) : (
                    <>
                      <Row className="mt-2">
                        <Col md={6}>
                          <Form.Group className="mb-3">
                            <Form.Label>Host</Form.Label>
                            <Form.Control
                              type="text"
                              placeholder="localhost"
                              name="Stat.Couchbase.Host"
                              value={values.Stat.Couchbase.Host}
                              onChange={handleChange}
                            />
                          </Form.Group>
                        </Col>
                        <Col md={6}>
                          <Form.Group className="mb-3">
                            <Form.Label>Bucket Name</Form.Label>
                            <Form.Control
                              type="text"
                              placeholder="encloud"
                              name="Stat.Couchbase.Bucket.Name"
                              value={values.Stat.Couchbase.Bucket.Name}
                              onChange={handleChange}
                            />
                          </Form.Group>
                        </Col>
                      </Row>
                      <Row>
                        <Col md={6}>
                          <Form.Group className="mb-3">
                            <Form.Label>Username</Form.Label>
                            <Form.Control
                              type="text"
                              placeholder="Administrator"
                              name="Stat.Couchbase.Username"
                              value={values.Stat.Couchbase.Username}
                              onChange={handleChange}
                            />
                          </Form.Group>
                        </Col>
                        <Col md={6}>
                          <Form.Group className="mb-3">
                            <Form.Label>Password</Form.Label>
                            <Form.Control
                              type="password"
                              placeholder="Encloud@2022"
                              name="Stat.Couchbase.Password"
                              value={values.Stat.Couchbase.Password}
                              onChange={handleChange}
                            />
                          </Form.Group>
                        </Col>
                      </Row>
                      <Row>
                        <Col md={6}>
                          <Form.Group className="mb-3">
                            <Form.Label>Scope Name</Form.Label>
                            <Form.Control
                              type="text"
                              placeholder="file"
                              name="Stat.Couchbase.Bucket.Scope"
                              value={values.Stat.Couchbase.Bucket.Scope}
                              onChange={handleChange}
                            />
                          </Form.Group>
                        </Col>
                        <Col md={6}>
                          <Form.Group className="mb-3">
                            <Form.Label>Collection Name</Form.Label>
                            <Form.Control
                              type="text"
                              placeholder="metadata"
                              name="Stat.Couchbase.Bucket.Collection"
                              value={values.Stat.Couchbase.Bucket.Collection}
                              onChange={handleChange}
                            />
                          </Form.Group>
                        </Col>
                      </Row>
                    </>
                  )}
                  <Row>
                    <Col md={12} className="text-left">
                      <ColoredBtn
                        className={`step-button ml-2 ${
                          loading ? "loadingStatus" : ""
                        }`}
                        disabled={loading}
                        onClick={handleSubmit}
                      >
                        {loading ? (
                          <div>
                            <ClipLoader
                              color="#ffffff"
                              loading={loading}
                              cssOverride={override}
                              size={30}
                              aria-label="Loading Spinner"
                              data-testid="loader"
                            />
                            <span className="loadingText">Submitting</span>
                          </div>
                        ) : (
                          "Submit"
                        )}
                      </ColoredBtn>
                      <ColoredBtn
                        className={`step-button ml-2 ${
                          restoreDefaultLoading ? "loadingStatus" : ""
                        }`}
                        style={{ marginLeft: 4 }}
                        disabled={restoreDefaultLoading}
                        onClick={restoreDefault}
                      >
                        {loading ? (
                          <div>
                            <ClipLoader
                              color="#ffffff"
                              loading={restoreDefaultLoading}
                              cssOverride={override}
                              size={30}
                              aria-label="Loading Spinner"
                              data-testid="loader"
                            />
                            <span className="loadingText">Submitting</span>
                          </div>
                        ) : (
                          "Restore Default"
                        )}
                      </ColoredBtn>
                    </Col>
                  </Row>
                </Form>
              )}
            </Formik>
          </KeyBoxedContent>
        </Card.Body>
      </Card>
    </>
  );
};

export default ConfigurationPage;
