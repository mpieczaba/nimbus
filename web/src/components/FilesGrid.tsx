import styled from "styled-components";

export const FilesGrid = styled.section`
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 2fr));
  column-gap: 1rem;
  grid-template-rows: 1fr auto;
  row-gap: 1rem;
  width: 100%;
  padding: 1rem;
`;
