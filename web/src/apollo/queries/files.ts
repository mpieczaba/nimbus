import { gql } from "@apollo/client";

export default gql`
  query {
    files {
      edges {
        node {
          id
          name
          size
          url
          updatedAt
        }
      }
    }
  }
`;

export interface FilesData {
  files: {
    edges: [
      {
        cursor: string;
        node: {
          id: string;
          name: string;
          size: number;
          url: string;
          updatedAt: Date;
        };
      }
    ];
  };
}
