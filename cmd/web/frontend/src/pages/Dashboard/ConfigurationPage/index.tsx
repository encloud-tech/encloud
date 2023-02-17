import * as formik from "formik";
import { Button, Card, Col, Form, Row } from "react-bootstrap";
import { KeyBoxedContent, KeyPairsSection } from "./styles";
import { PageHeader } from "./../../../components/layouts/styles";
import Select, { StylesConfig } from "react-select";
import { useState } from "react";

const storageTypeOptions = [
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
  const { Formik } = formik;

  const save = async (data: any) => {
    console.log(data);
  };

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
            <Formik
              onSubmit={save}
              initialValues={{
                estuary: {
                  base_api_url: "",
                  token: "",
                },
                email: {
                  server: "",
                  port: "",
                  username: "",
                  password: "",
                  from: "",
                },
                stat: {
                  kekType: kekTypeOptions[0],
                  storageType: storageTypeOptions[0],
                  badgerdb: {
                    path: "",
                  },
                  couchbase: {
                    host: "",
                    username: "",
                    password: "",
                    bucket: {
                      name: "",
                      scope: "",
                      collection: "",
                    },
                  },
                },
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
                          name="estuary.base_api_url"
                          placeholder="https://api.estuary.tech"
                          value={values.estuary.base_api_url}
                          onChange={handleChange}
                        />
                      </Form.Group>
                    </Col>
                    <Col md={6}>
                      <Form.Group className="mb-3">
                        <Form.Label>Token</Form.Label>
                        <Form.Control
                          type="text"
                          name="estuary.token"
                          placeholder="EST6315eb22-XXXX-XXXX-XXXX-1acb4a954070ARY"
                          value={values.estuary.token}
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
                        <Form.Label>Server</Form.Label>
                        <Form.Control
                          type="text"
                          placeholder="smtp.mailtrap.io"
                          name="email.server"
                          value={values.email.server}
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
                          name="email.port"
                          value={values.email.port}
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
                          name="email.username"
                          value={values.email.username}
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
                          name="email.password"
                          value={values.email.password}
                          onChange={handleChange}
                        />
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
                          name="email.from"
                          value={values.email.from}
                          onChange={handleChange}
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
                          name="stat.kekType"
                          value={values.stat.kekType}
                          onChange={(newVal) => {
                            setFieldValue("stat.kekType", newVal);
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
                          name="stat.storageType"
                          value={values.stat.storageType}
                          onChange={(newVal) => {
                            setFieldValue("stat.storageType", newVal);
                          }}
                        />
                      </Form.Group>
                    </Col>
                  </Row>
                  {values.stat.storageType.value === "badgerdb" ? (
                    <Row className="mt-2">
                      <Col md={6}>
                        <Form.Group className="mb-3">
                          <Form.Label>Path</Form.Label>
                          <Form.Control
                            type="text"
                            placeholder="badger.db"
                            name="stat.badgerdb.path"
                            value={values.stat.badgerdb.path}
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
                              name="stat.couchbase.host"
                              value={values.stat.couchbase.host}
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
                              name="stat.couchbase.bucket.name"
                              value={values.stat.couchbase.bucket.name}
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
                              name="stat.couchbase.username"
                              value={values.stat.couchbase.username}
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
                              name="stat.couchbase.password"
                              value={values.stat.couchbase.password}
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
                              name="stat.couchbase.bucket.scope"
                              value={values.stat.couchbase.bucket.scope}
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
                              name="stat.couchbase.bucket.collection"
                              value={values.stat.couchbase.bucket.collection}
                              onChange={handleChange}
                            />
                          </Form.Group>
                        </Col>
                      </Row>
                    </>
                  )}
                  <Row>
                    <Col md={12} className="text-left">
                      <Button type="submit">Submit</Button>
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
