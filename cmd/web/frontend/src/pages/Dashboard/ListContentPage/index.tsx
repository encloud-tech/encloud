import * as Yup from "yup";
import * as formik from "formik";
import { Card, Image, Button, Modal, Form, Col } from "react-bootstrap";
import { PageHeader } from "../../../components/layouts/styles";
import { SectionBox } from "./styles";
import { useEffect, useState } from "react";
import {
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from "@tanstack/react-table";
import { Table as BTable } from "react-bootstrap";

import menuIcon from "../../../assets/images/menu.png";
import { List, Share } from "../../../../wailsjs/go/main/App";
import { readKey } from "../../../services/localStorage.service";
import { types } from "../../../../wailsjs/go/models";
import { Link } from "react-router-dom";
import { toast } from "react-toastify";

const columnHelper = createColumnHelper<types.FileMetadata>();

const ListContentPage = () => {
  const [data, setData] = useState<types.FileMetadata[]>([]);
  const [open, setOpen] = useState(false);
  const [selected, setSelected] = useState<types.FileMetadata>();
  const [shareLoading, setShareLoading] = useState(false);

  const { Formik } = formik;

  const schema = Yup.object().shape({
    email: Yup.string().required("Please enter email"),
  });

  const columns = [
    columnHelper.accessor("name", {
      cell: (info) => info.getValue(),
      header: () => <span>File Name</span>,
    }),
    columnHelper.accessor("size", {
      cell: (info) => info.getValue(),
      header: () => <span>File Size</span>,
    }),
    columnHelper.accessor("cid", {
      cell: (info) => info.getValue(),
      header: () => <span>CID</span>,
    }),
    columnHelper.accessor("uuid", {
      header: () => "Actions",
      cell: (info) => (
        <>
          <Link
            to={`/retrieve/${info.getValue()}`}
            className="btn btn-primary list-button"
            style={{ marginRight: 5 }}
            state={{ metadata: info.row.original }}
          >
            Retrieve
          </Link>
          <Button
            className="list-button"
            onClick={() => {
              setSelected(info.row.original);
              setOpen(true);
            }}
          >
            Share
          </Button>
        </>
      ),
    }),
  ];

  useEffect(() => {
    const fetchData = async () => {
      const response = await List(readKey()?.PublicKey);

      if (response.Data) {
        setData(response.Data);
      }
    };

    fetchData();
  }, []);

  const share = async (data: any) => {
    setShareLoading(true);
    if (selected) {
      try {
        Share(
          selected?.uuid,
          readKey().PublicKey,
          readKey().PrivateKey,
          data.email
        )
          .then((result) => {
            if (result && result.Status == "success") {
              setShareLoading(false);
              setOpen(false);
              toast.success("Document shared successfully.", {
                position: toast.POSITION.TOP_RIGHT,
              });
            } else {
              setOpen(false);
              setShareLoading(false);
              toast.error(result.Message, {
                position: toast.POSITION.TOP_RIGHT,
              });
            }
          })
          .catch((err: any) => {
            setOpen(false);
            setShareLoading(false);
            toast.error("Something went wrong!.Please retry", {
              position: toast.POSITION.TOP_RIGHT,
            });
          });
      } catch (err) {
        setOpen(false);
        setShareLoading(false);
        toast.error("Something went wrong!.Please retry", {
          position: toast.POSITION.TOP_RIGHT,
        });
      }
    }
  };

  const table = useReactTable({
    data,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <>
      <PageHeader>
        <h2>
          <Image className="titleIcon" src={menuIcon} />
          <span>Content List</span>
        </h2>
      </PageHeader>
      <SectionBox>
        <Card>
          <Card.Body>
            <BTable bordered hover responsive>
              <thead>
                {table.getHeaderGroups().map((headerGroup) => (
                  <tr key={headerGroup.id}>
                    {headerGroup.headers.map((header) => (
                      <th key={header.id}>
                        {header.isPlaceholder
                          ? null
                          : flexRender(
                              header.column.columnDef.header,
                              header.getContext()
                            )}
                      </th>
                    ))}
                  </tr>
                ))}
              </thead>
              <tbody>
                {table.getRowModel().rows.map((row) => (
                  <tr key={row.id}>
                    {row.getVisibleCells().map((cell) => (
                      <td key={cell.id}>
                        {flexRender(
                          cell.column.columnDef.cell,
                          cell.getContext()
                        )}
                      </td>
                    ))}
                  </tr>
                ))}
              </tbody>
            </BTable>
          </Card.Body>
        </Card>
      </SectionBox>
      <Modal show={open} onHide={() => setOpen(false)}>
        <Modal.Header closeButton>
          <Modal.Title>Share Content</Modal.Title>
        </Modal.Header>
        <Formik
          validationSchema={schema}
          onSubmit={share}
          initialValues={{
            email: "",
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
          }) => (
            <Form noValidate onSubmit={handleSubmit}>
              <Modal.Body>
                <Form.Control
                  type="text"
                  placeholder="Email"
                  aria-label="Email"
                  name="email"
                  value={values.email}
                  onChange={handleChange}
                  isInvalid={!!errors.email}
                />
                <span
                  className="invalid-feedback"
                  style={{ color: "red", textAlign: "left" }}
                >
                  {errors.email}
                </span>
              </Modal.Body>
              <Modal.Footer>
                <Button variant="secondary" onClick={() => setOpen(false)}>
                  Close
                </Button>
                <Button
                  type="submit"
                  variant="primary"
                  style={{ padding: 7, fontSize: "inherit" }}
                >
                  Share
                </Button>
              </Modal.Footer>
            </Form>
          )}
        </Formik>
      </Modal>
    </>
  );
};

export default ListContentPage;
