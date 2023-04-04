import * as Yup from "yup";
import * as formik from "formik";
import {
  Button,
  Card,
  Col,
  Form,
  Image,
  OverlayTrigger,
  Row,
  Tooltip,
} from "react-bootstrap";
import { KeyBoxedContent, KeyPairsSection } from "./styles";
import { PageHeader } from "../../../components/layouts/styles";

// Images
import shareIcon from "../../../assets/images/share.png";
import {
  RetrieveSharedContent,
  SelectDirectory,
  SelectFile,
} from "../../../../wailsjs/go/main/App";
import { toast } from "react-toastify";
import { CSSProperties, useState } from "react";
import ClipLoader from "react-spinners/ClipLoader";
import { ColoredBtn } from "./styles";
import Select, { StylesConfig } from "react-select";

const override: CSSProperties = {
  display: "block",
  margin: "0 auto",
  borderColor: "white",
};

const dekTypeOptions = [
  { value: "aes", label: "AES 256 GCM" },
  { value: "chacha20", label: "ChaCha20-Poly1305" },
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
      <strong>AES 256 GCM</strong> - GCM throughput rates for state-of-the-art,
      high-speed communication channels can be achieved with inexpensive
      hardware resources. GCM is limited to encrypting 64 GiB of plain text.
    </p>
    <p>
      <strong>ChaCha20-Poly1035</strong> - ChaCha20-Poly1305 is an AEAD
      algorithm, that combines the ChaCha20 stream cipher with the Poly1305
      message authentication code. Without hardware acceleration, it is usually
      faster than AES-GCM.
    </p>
  </Tooltip>
);

const RetrieveSharedContentPage = () => {
  const [loading, setLoading] = useState(false);
  const [selectedDirectory, setSelectedDirectory] = useState<string>();
  const [dekPath, setDekPath] = useState<string>();
  const { Formik } = formik;

  const schema = Yup.object().shape({
    cid: Yup.string().required("Please enter cid"),
    fileName: Yup.string().required("Please enter download file name"),
  });

  const getFilePath = async (evt: any) => {
    evt.preventDefault();
    try {
      SelectFile()
        .then((result: any) => {
          var dt = new DataTransfer();
          dt.items.add(new File([], result));
          evt.target.files = dt.files;
          setDekPath(result);
        })
        .catch((err: any) => {
          console.error(err);
        });
    } catch (err) {
      console.error(err);
    }
  };

  const getDirectoryPath = async (evt: any) => {
    evt.preventDefault();
    try {
      SelectDirectory()
        .then((result: any) => {
          var dt = new DataTransfer();
          dt.items.add(new File([], result));
          evt.target.files = dt.files;
          setSelectedDirectory(result);
        })
        .catch((err: any) => {
          console.error(err);
        });
    } catch (err) {
      console.error(err);
    }
  };

  const getSharedContent = (data: any) => {
    setLoading(true);
    try {
      if (selectedDirectory && dekPath) {
        RetrieveSharedContent(
          dekPath,
          data.dekType.value,
          data.cid,
          data.fileName,
          selectedDirectory
        )
          .then((result: any) => {
            if (result && result.Status == "success") {
              setLoading(false);
              toast.success("Document downloaded successfully.", {
                position: toast.POSITION.TOP_RIGHT,
              });
            } else {
              setLoading(false);
              toast.error("Something went wrong!.Please retry", {
                position: toast.POSITION.TOP_RIGHT,
              });
            }
          })
          .catch((err: any) => {
            setLoading(false);
            toast.error("Something went wrong!.Please retry", {
              position: toast.POSITION.TOP_RIGHT,
            });
          });
      }
    } catch (err) {
      setLoading(false);
      toast.error("Something went wrong!.Please retry", {
        position: toast.POSITION.TOP_RIGHT,
      });
    }
  };

  return (
    <>
      <PageHeader>
        <h2>
          <Image className="titleIcon" src={shareIcon} />
          <span>Retrieve Data Received</span>
        </h2>
      </PageHeader>
      <Card className="mt-4">
        <Card.Body>
          <KeyBoxedContent>
            <Formik
              validationSchema={schema}
              onSubmit={getSharedContent}
              initialValues={{
                cid: "",
                dekType: dekTypeOptions[0],
                fileName: "",
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
                    <Col md={12} lg={6}>
                      <KeyPairsSection>
                        <h3>Retrieve Data Received</h3>
                      </KeyPairsSection>
                    </Col>
                  </Row>
                  <Row className="mt-4">
                    <Col md={8}>
                      <Form.Group className="mb-3">
                        <Form.Label>CID</Form.Label>
                        <Form.Control
                          type="text"
                          placeholder="CID"
                          name="cid"
                          value={values.cid}
                          onChange={handleChange}
                          isInvalid={!!errors.cid}
                        />
                        <span
                          className="invalid-feedback"
                          style={{ color: "red", textAlign: "left" }}
                        >
                          {errors.cid}
                        </span>
                      </Form.Group>
                    </Col>
                    <Col md={4}></Col>
                  </Row>
                  <Row className="mt-4">
                    <Col md={8}>
                      <Form.Group className="mb-3">
                        <Form.Label>
                          DEK Type
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
                          name="dekType"
                          className="dek-type-select"
                          styles={colourStyles}
                          options={dekTypeOptions}
                          value={values.dekType}
                          onChange={(newVal) => {
                            setFieldValue("dekType", newVal);
                          }}
                        />
                      </Form.Group>
                    </Col>
                    <Col md={4}></Col>
                  </Row>
                  <Row className="mt-4">
                    <Col md={8}>
                      <Form.Group className="mb-3">
                        <Form.Label>DEK Path</Form.Label>
                        <Form.Control
                          type="file"
                          name="decryptedDekPath"
                          onClick={getFilePath}
                        />
                      </Form.Group>
                    </Col>
                    <Col md={4}></Col>
                  </Row>
                  <Row className="mt-4">
                    <Col md={8}>
                      <Form.Group className="mb-3">
                        <Form.Label>Download File Name</Form.Label>
                        <Form.Control
                          type="text"
                          placeholder="sample.csv"
                          name="fileName"
                          value={values.fileName}
                          onChange={handleChange}
                          isInvalid={!!errors.fileName}
                        />
                        <span
                          className="invalid-feedback"
                          style={{ color: "red", textAlign: "left" }}
                        >
                          {errors.fileName}
                        </span>
                      </Form.Group>
                    </Col>
                    <Col md={4}></Col>
                  </Row>
                  <Row className="mt-4">
                    <Col md={8}>
                      <Form.Group className="mb-3">
                        <Form.Label>Download File Path</Form.Label>
                        <Form.Control
                          type="file"
                          name="retrievalFileStoragePath"
                          onClick={getDirectoryPath}
                        />
                      </Form.Group>
                    </Col>
                    <Col md={4}></Col>
                  </Row>
                  <Row>
                    <Col md={8} className="text-left">
                      <ColoredBtn
                        className={`step-button ml-2 ${
                          loading ? "loadingStatus" : ""
                        }`}
                        disabled={loading || !selectedDirectory || !dekPath}
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
                            <span className="loadingText">Downloading</span>
                          </div>
                        ) : (
                          "Get Shared Content"
                        )}
                      </ColoredBtn>
                    </Col>
                    <Col md={4}></Col>
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

export default RetrieveSharedContentPage;
