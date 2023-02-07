import { Badge, Button, Card, Col, Form, Image, Row } from "react-bootstrap";
import Dropzone from "react-dropzone";
import Select, { StylesConfig } from "react-select";

import { DragArea, FileUploadBar } from "./styles";
import { PageHeader } from "./../../../components/layouts/styles";

import dsupload1Img from "../../../assets/images/ds-upload-1.png";

const options = [
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
                    className="dek-type-select"
                    styles={colourStyles}
                    options={options}
                    defaultValue={{ value: "AES", label: "AES 256 GCM" }}
                  />
                </Form.Group>
              </Col>
              <Col md={4}></Col>
            </Row>
            <Row>
              <Col md={8}>
                <Form.Group className="mb-3">
                  <Form.Label>
                    Chunk Sizes <Badge bg="success">Preminum</Badge>{" "}
                  </Form.Label>
                  <Form.Control type="text" placeholder="Enter Chunk Sizes" />
                </Form.Group>
              </Col>
              <Col md={4}></Col>
            </Row>
            <Row>
              <Col md={8}>
                <Dropzone>
                  {({ getRootProps, getInputProps }) => (
                    <DragArea {...getRootProps()}>
                      <span className="choose-file-button">Choose files</span>
                      <span className="file-message">
                        or drag and drop files here
                      </span>
                      <input className="file-input" {...getInputProps()} />
                    </DragArea>
                  )}
                </Dropzone>
              </Col>
              <Col md={4}></Col>
            </Row>
            <Row>
              <Col md={8} className="text-right">
                <Button>Upload</Button>
              </Col>
              <Col md={4}></Col>
            </Row>
          </Form>
          <Row className="mt-3">
            <Col md={8}>
              <FileUploadBar now={60} label={`60%`} />
            </Col>
            <Col md={4}></Col>
          </Row>
        </Card.Body>
      </Card>
    </>
  );
};

export default UploadContent;
