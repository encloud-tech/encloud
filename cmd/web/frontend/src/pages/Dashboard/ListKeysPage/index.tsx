import { Card, Image } from "react-bootstrap";
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
import { ListKeys } from "../../../../wailsjs/go/main/App";
import { types } from "../../../../wailsjs/go/models";
import { Link } from "react-router-dom";

const columnHelper = createColumnHelper<types.FetchKeysResponse>();

const ListKeysPage = () => {
  const [data, setData] = useState<types.FetchKeysResponse[]>([]);

  const columns = [
    columnHelper.accessor("publicKey", {
      cell: (info) => info.getValue(),
      header: () => <span>Key</span>,
    }),
    columnHelper.accessor("files", {
      cell: (info) => info.getValue(),
      header: () => <span>No of Files</span>,
    }),
    columnHelper.accessor("publicKey", {
      header: () => "Actions",
      cell: (info) => (
        <Link
          to={`/list`}
          className="btn btn-primary list-button"
          style={{ marginRight: 5 }}
          state={{ metadata: info.row.original }}
        >
          View
        </Link>
      ),
    }),
  ];

  useEffect(() => {
    const fetchData = async () => {
      const response = await ListKeys();

      if (response.Data) {
        setData(response.Data);
      }
    };

    fetchData();
  }, []);

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
    </>
  );
};

export default ListKeysPage;
