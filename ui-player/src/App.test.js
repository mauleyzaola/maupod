import { configure, shallow }  from 'enzyme'
import { expect } from 'chai'
import Adapter from 'enzyme-adapter-react-16'
import React from 'react'
import { TrackListRow} from "./components/TrackList"

configure({ adapter: new Adapter() })

describe('sample test', () => {
  const row = {
    format: 'FLAC',
  }
  const wrapper = shallow(<TrackListRow i={1} row={row} />)

  it('should render format', () => {
    const expectedDiv = <td>FLAC</td>
    const actualValue = wrapper.containsMatchingElement(expectedDiv)
    expect(actualValue).to.equal(true)
  })
})