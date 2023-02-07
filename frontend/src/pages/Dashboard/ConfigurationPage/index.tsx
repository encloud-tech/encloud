import { Button, Card, Col, Form, Row } from "react-bootstrap";
import { KeyBoxedContent, KeyPairsSection } from "./styles";
import { PageHeader } from "./../../../components/layouts/styles";
import Select, { StylesConfig } from "react-select";
import { useState } from "react";

const options = [
  { value: "badgerdb", label: "Badger DB" },
  { value: "couchbase", label: "Couchbase DB" },
];

const kekTypeOptions = [
  { value: "rsa", label: "RSA" },
  { value: "ecies", label: "ECIES" },
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

const ConfigurationPage = () => {
  const [storageType, setStorageType] = useState(options[0]);

  return (
    <>
      <PageHeader>
        <h2>
          <span>Configuration</span>
        </h2>
      </PageHeader>
      <Card className="mt-4">
        <Card.Body>
          <KeyBoxedContent>
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
                    placeholder="https://api.estuary.tech"
                  />
                </Form.Group>
              </Col>
              <Col md={6}>
                <Form.Group className="mb-3">
                  <Form.Label>Token</Form.Label>
                  <Form.Control
                    type="text"
                    placeholder="EST6315eb22-XXXX-XXXX-XXXX-1acb4a954070ARY"
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
                  <Form.Label>Server</Form.Label>
                  <Form.Control type="text" placeholder="smtp.mailtrap.io" />
                </Form.Group>
              </Col>
              <Col md={6}>
                <Form.Group className="mb-3">
                  <Form.Label>Port</Form.Label>
                  <Form.Control type="text" placeholder="2525" />
                </Form.Group>
              </Col>
            </Row>
            <Row>
              <Col md={6}>
                <Form.Group className="mb-3">
                  <Form.Label>Username</Form.Label>
                  <Form.Control type="text" placeholder="ac984e52bfd98h" />
                </Form.Group>
              </Col>
              <Col md={6}>
                <Form.Group className="mb-3">
                  <Form.Label>Password</Form.Label>
                  <Form.Control type="password" placeholder="861b495c076987" />
                </Form.Group>
              </Col>
            </Row>
            <Row>
              <Col md={6}>
                <Form.Group className="mb-3">
                  <Form.Label>From Address</Form.Label>
                  <Form.Control
                    type="text"
                    placeholder="noreply@encloud.tech"
                  />
                </Form.Group>
              </Col>
              <Col md={6}></Col>
            </Row>
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
                  <Form.Label>Kek Type</Form.Label>
                  <Select
                    className="dek-type-select"
                    styles={colourStyles}
                    options={kekTypeOptions}
                    defaultValue={kekTypeOptions[0]}
                  />
                </Form.Group>
              </Col>
              <Col md={6}>
                <Form.Group className="mb-3">
                  <Form.Label>Storage Type</Form.Label>
                  <Select
                    name="storageType"
                    className="dek-type-select"
                    styles={colourStyles}
                    options={options}
                    value={storageType}
                    onChange={(val: any) => {
                      setStorageType(val);
                    }}
                  />
                </Form.Group>
              </Col>
            </Row>
            {storageType.value === "badgerdb" ? (
              <Row className="mt-2">
                <Col md={6}>
                  <Form.Group className="mb-3">
                    <Form.Label>Path</Form.Label>
                    <Form.Control type="text" placeholder="badger.db" />
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
                      <Form.Control type="text" placeholder="localhost" />
                    </Form.Group>
                  </Col>
                  <Col md={6}>
                    <Form.Group className="mb-3">
                      <Form.Label>Bucket Name</Form.Label>
                      <Form.Control type="text" placeholder="encloud" />
                    </Form.Group>
                  </Col>
                </Row>
                <Row>
                  <Col md={6}>
                    <Form.Group className="mb-3">
                      <Form.Label>Username</Form.Label>
                      <Form.Control type="text" placeholder="Administrator" />
                    </Form.Group>
                  </Col>
                  <Col md={6}>
                    <Form.Group className="mb-3">
                      <Form.Label>Password</Form.Label>
                      <Form.Control
                        type="password"
                        placeholder="Encloud@2022"
                      />
                    </Form.Group>
                  </Col>
                </Row>
                <Row>
                  <Col md={6}>
                    <Form.Group className="mb-3">
                      <Form.Label>Scope Name</Form.Label>
                      <Form.Control type="text" placeholder="file" />
                    </Form.Group>
                  </Col>
                  <Col md={6}>
                    <Form.Group className="mb-3">
                      <Form.Label>Collection Name</Form.Label>
                      <Form.Control type="text" placeholder="metadata" />
                    </Form.Group>
                  </Col>
                </Row>
              </>
            )}
            <Row>
              <Col md={12} className="text-left">
                <Button>Submit</Button>
              </Col>
            </Row>
          </KeyBoxedContent>
        </Card.Body>
      </Card>
    </>
  );
};

export default ConfigurationPage;
