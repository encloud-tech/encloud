import {
  Accordion,
  Badge,
  Button,
  Card,
  Col,
  Form,
  Image,
  OverlayTrigger,
  Row,
  Tooltip,
} from "react-bootstrap";
import Select, { StylesConfig } from "react-select";

import { PageHeader } from "./../../../components/layouts/styles";

import uploadIcon from "../../../assets/images/upload.png";
import { CSSProperties, useState } from "react";
import {
  GenerateKeyPair,
  SelectFile,
  Upload,
} from "../../../../wailsjs/go/main/App";
import {
  persistKekType,
  persistKey,
  readKekType,
  readKey,
} from "../../../services/localStorage.service";
import ClipLoader from "react-spinners/ClipLoader";
import { ColoredBtn } from "./styles";
import { toast } from "react-toastify";

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

const UploadContent = () => {
  const [uploadLoading, setUploadLoading] = useState(false);
  const [dekType, setDekType] = useState(dekTypeOptions[0]);
  const [filePath, setFilePath] = useState<string>();

  const getFilePath = async (evt: any) => {
    evt.preventDefault();
    try {
      SelectFile()
        .then((result: any) => {
          var dt = new DataTransfer();
          dt.items.add(new File([], result));
          evt.target.files = dt.files;
          setFilePath(result);
        })
        .catch((err: any) => {
          console.error(err);
        });
    } catch (err) {
      console.error(err);
    }
  };

  const doUpload = async () => {
    setUploadLoading(true);
    if (filePath) {
      try {
        let kekType = readKekType();
        if (!kekType) {
          kekType = "rsa";
          const response = await GenerateKeyPair("rsa");
          persistKey(response.Data);
          persistKekType("rsa");
        }
        Upload(filePath, readKekType(), dekType.value, readKey().PublicKey)
          .then((result: any) => {
            if (result && result.Status == "success") {
              setFilePath("");
              setUploadLoading(false);
              toast.success("Document uploaded successfully.", {
                position: toast.POSITION.TOP_RIGHT,
              });
            } else {
              setFilePath("");
              setUploadLoading(false);
              toast.error("Something went wrong!.Please retry", {
                position: toast.POSITION.TOP_RIGHT,
              });
            }
          })
          .catch((err: any) => {
            setFilePath("");
            setUploadLoading(false);
            toast.error("Something went wrong!.Please retry", {
              position: toast.POSITION.TOP_RIGHT,
            });
          });
      } catch (err) {
        setFilePath("");
        setUploadLoading(false);
        toast.error("Something went wrong!.Please retry", {
          position: toast.POSITION.TOP_RIGHT,
        });
      }
    }
  };

  return (
    <>
      <PageHeader>
        <h2>
          <Image className="titleIcon" src={uploadIcon} />
          <span>Upload</span>
        </h2>
      </PageHeader>

      <Card className="mb-3">
        <Card.Body className="p-3">
          <Row className="mb-3">
            <Col>
              <span className="fw-bold">
                {uploadLoading ? "Upload is in progress..." : ""}
              </span>
            </Col>
          </Row>
          <Form className="box">
            <Row>
              <Col md={8}>
                <Form.Group className="mb-3">
                  <Form.Label>Document</Form.Label>
                  <Form.Control type="file" onClick={getFilePath} />
                </Form.Group>
              </Col>
              <Col md={4}></Col>
            </Row>
            <Row>
              <Col md={8}>
                <Accordion>
                  <Accordion.Item eventKey="0">
                    <Accordion.Header>Advance Options</Accordion.Header>
                    <Accordion.Body>
                      <Row>
                        <Col md={12}>
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
                              value={dekType}
                              onChange={(newVal: any) => {
                                setDekType(newVal);
                              }}
                            />
                          </Form.Group>
                        </Col>
                      </Row>
                      <Row>
                        <Col md={12}>
                          <Form.Group className="mb-3">
                            <Form.Label>
                              Chunk Sizes <Badge bg="success">Premium</Badge>{" "}
                            </Form.Label>
                            <Form.Control
                              disabled={true}
                              readOnly={true}
                              type="text"
                              placeholder="Enter Chunk Sizes"
                            />
                          </Form.Group>
                        </Col>
                      </Row>
                    </Accordion.Body>
                  </Accordion.Item>
                </Accordion>
              </Col>
              <Col md={4}></Col>
            </Row>
            <Row className="mt-3">
              <Col md={8} className="text-right">
                <ColoredBtn
                  className={`step-button ml-2 ${
                    uploadLoading ? "loadingStatus" : ""
                  }`}
                  disabled={uploadLoading || !filePath}
                  onClick={doUpload}
                >
                  {uploadLoading ? (
                    <div>
                      <ClipLoader
                        color="#ffffff"
                        loading={uploadLoading}
                        cssOverride={override}
                        size={30}
                        aria-label="Loading Spinner"
                        data-testid="loader"
                      />
                      <span className="loadingText">Uploading</span>
                    </div>
                  ) : (
                    "Upload"
                  )}
                </ColoredBtn>
              </Col>
              <Col md={4}></Col>
            </Row>
          </Form>
        </Card.Body>
      </Card>
    </>
  );
};

export default UploadContent;
