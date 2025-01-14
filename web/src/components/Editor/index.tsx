import ReactQuill from 'react-quill-new'
import 'react-quill-new/dist/quill.snow.css'

type EditorProps = {
  label: string
  value: string
  setValue: (value: string) => void
}
const Editor: React.FC<EditorProps> = ({ label, value, setValue }) => {
  return (
    <div className="relative mt-2">
      <ReactQuill
        value={value}
        id="editor"
        onChange={setValue}
        theme="snow"
        modules={{
          toolbar: [
            ['bold', 'italic', 'underline', 'strike'], // toggled buttons
            [
              'blockquote',
              // 'code-block', NOTE: do not work
            ],
            [
              'link',
              'image',
              'video',
              // 'formula', // NOTE: do not work
            ],

            [{ header: 1 }, { header: 2 }], // custom button values
            [{ list: 'ordered' }, { list: 'bullet' }, { list: 'check' }],
            // [{ script: 'sub' }, { script: 'super' }], // superscript/subscript: NOTE: do not work
            [{ indent: '-1' }, { indent: '+1' }], // outdent/indent
            // [{ direction: 'rtl' }], // text direction // NOTE: do not work

            [{ size: ['small', false, 'large', 'huge'] }], // custom dropdown
            [{ header: [1, 2, 3, 4, 5, 6, false] }],

            // [{ color: [] }, { background: [] }], // dropdown with defaults from theme NOTE: do not work
            [{ font: [] }],
            // [{ align: [] }], NOTE: do not work

            ['clean'], // remove formatting button
          ],
        }}
        formats={[
          'header',
          'font',
          'size',
          'bold',
          'italic',
          'underline',
          'strike',
          'blockquote',
          'list',
          'indent',
          'link',
          'image',
          'video',
        ]}
      />

      <label className="absolute left-3 -top-2 bg-white px-1 text-xs text-gray-700">
        {label}
      </label>
    </div>
  )
}

export default Editor
