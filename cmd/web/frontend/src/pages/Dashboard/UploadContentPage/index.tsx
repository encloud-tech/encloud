import { Badge, Card, Col, Form, Image, Row } from "react-bootstrap";
import Select, { StylesConfig } from "react-select";

import { PageHeader } from "./../../../components/layouts/styles";

import dsupload1Img from "../../../assets/images/ds-upload-1.png";
import { CSSProperties, useState } from "react";
import { SelectFile, Upload } from "../../../../wailsjs/go/main/App";
import { readKey } from "../../../services/localStorage.service";
import { useNavigate } from "react-router-dom";
import ClipLoader from "react-spinners/ClipLoader";
import { ColoredBtn } from "./styles";

const override: CSSProperties = {
  display: "block",
  margin: "0 auto",
  borderColor: "white",
};

const dekTypeOptions = [
  { value: "AES", label: "AES 256 GCM" },
  { value: "Chacha20", label: "ChaCha20-Poly1305" },
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

const UploadContent = () => {
  const [uploadLoading, setUploadLoading] = useState(false);
  const [dekType, setDekType] = useState(dekTypeOptions[0]);
  const [filePath, setFilePath] = useState<string>();

  const navigate = useNavigate();

  const getFilePath = async (evt: any) => {
    evt.preventDefault();
    try {
      SelectFile()
        .then((result: any) => {
          setFilePath(result);
        })
        .catch((err: any) => {
          console.error(err);
        });
    } catch (err) {
      console.error(err);
    }
  };

  const doUpload = () => {
    setUploadLoading(true);
    if (filePath) {
      try {
        Upload(filePath, "rsa", dekType.value, readKey().PublicKey)
          .then((result: any) => {
            if (result && result.Data) {
              navigate("/list");
              setUploadLoading(false);
            }
          })
          .catch((err: any) => {
            console.error(err);
            setUploadLoading(false);
          });
      } catch (err) {
        console.error(err);
        setUploadLoading(false);
      }
    }
  };

  return (
    <>
      <PageHeader>
        <h2>
          <Image className="titleIcon" src={dsupload1Img} />
          <span>Upload</span>
        </h2>
      </PageHeader>

      <Card className="mb-3">
        <Card.Body className="p-3">
          <Form className="box">
            <Row>
              <Col md={8}>
                <Form.Group className="mb-3">
                  <Form.Label>DEK Type</Form.Label>
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
              <Col md={4}></Col>
            </Row>
            <Row>
              <Col md={8}>
                <Form.Group className="mb-3">
                  <Form.Label>
                    Chunk Sizes <Badge bg="success">Premium</Badge>{" "}
                  </Form.Label>
                  <Form.Control type="text" placeholder="Enter Chunk Sizes" />
                </Form.Group>
              </Col>
              <Col md={4}></Col>
            </Row>
            <Row>
              <Col md={8}>
                <Form.Control type="file" onClick={getFilePath} />
              </Col>
              <Col md={4}></Col>
            </Row>
            <Row className="mt-3">
              <Col md={8} className="text-right">
                <ColoredBtn
                  className={`step-button ml-2 ${
                    uploadLoading ? "loadingStatus" : ""
                  }`}
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
