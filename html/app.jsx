class SearchBar extends React.Component {
  constructor(props) {
    super(props);

    this.handleInputChange = this.handleInputChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleInputChange(event) {
    this.props.onSearchInput(event.target.value);
  }

  handleSubmit(event) {
    this.props.onSearchSubmit(event);
  }

  render() {
    let btn =
      <ReactBootstrap.Button
        type="submit"
        onClick={this.handleSubmit}>
        Generate
      </ReactBootstrap.Button>
    if (this.props != null && this.props.appName != null) {
      btn =
        <DeleteAppButton
          onDeleteApp={this.props.onDeleteApp}
          appName={this.props.appName}
          />
    }
    return (
      <ReactBootstrap.Form horizontal>
        <ReactBootstrap.FormGroup>
          <ReactBootstrap.Row>
            <ReactBootstrap.Col xs={10}>
              <ReactBootstrap.FormControl
                autoFocus="true"
                type="text"
                placeholder="Feed me Prow URLs..."
                value={this.props.searchInput}
                onChange={this.handleInputChange}
              />
            </ReactBootstrap.Col>
            <ReactBootstrap.Col xs={2}>
            {btn}
            </ReactBootstrap.Col>
          </ReactBootstrap.Row>
        </ReactBootstrap.FormGroup>
      </ReactBootstrap.Form>
    );
  }
}

class DeleteAppButton extends React.Component {
  render() {
    return (
      <ReactBootstrap.Button variant="warning" onClick={this.props.onDeleteApp}>
      Delete {this.props.appName}
      </ReactBootstrap.Button>
    )
  }
}

class Message extends React.Component {
  render() {
    var lines = this.props.message.trim().split('\n');
    var variants = {
      "status": "info",
      "progress": "info",
      "failure": "danger",
      "done": "success"
    }
    switch (this.props.action) {
      case 'done':
      case 'failure':
      case 'status':
        return (
          <ReactBootstrap.Alert className="alert-small" variant={variants[this.props.action]}>
            {this.props.message}
          </ReactBootstrap.Alert>
        )
        break;
      case 'progress':
        return (
          <ReactBootstrap.Alert className="alert-small" variant={variants[this.props.action]}>
            <ReactBootstrap.Spinner animation="grow" size="sm" /><span>{this.props.message}</span>
          </ReactBootstrap.Alert>
        )
        break;
      case 'link':
        return (
          <ReactBootstrap.Alert className="alert-small" variant="primary">
            <ReactBootstrap.Alert.Link href={this.props.message}>{this.props.message}</ReactBootstrap.Alert.Link>
          </ReactBootstrap.Alert>
        )
        break;
      default:
        return (
          <span></span>
        )
    }
  }
}

class ResourceQuotaStatus extends React.Component {
  render() {
    if (this.props === null || this.props.resourceQuota === null) {
      return (
        <span></span>
      )
    }
    let used = this.props.resourceQuota.used;
    let hard = this.props.resourceQuota.hard;
    if (typeof used == "undefined" || typeof hard == "undefined") {
      return (
        <span></span>
      )
    }
    return (
      <div>
        <div>Current resource quota</div>
        <ReactBootstrap.ProgressBar now={used} max={hard}
          label={used + "/" +hard}/>
      </div>
    )
  }
}

class Status extends React.Component {
  render() {
    if (this.props.messages.length == 0) {
      return (
        <span></span>
      );
    } else {
      return(
        <div>
          {
            this.props.messages.map(item =>
              <Message
                action={item.action}
                message={item.message}
                onDeleteApp={this.props.onDeleteApp}
              />
            )
          }
        </div>
      );
    }
  }
}

class SearchForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      querySearch: '',
      searchInput: '',
      messages: [],
      appName: null,
      ws: null,
      resourceQuota: {
        used: 0,
        hard: 0,
      }
    };

    this.handleSearchInput = this.handleSearchInput.bind(this);
    this.handleSearchSubmit = this.handleSearchSubmit.bind(this);
    this.handleDeleteApp = this.handleDeleteApp.bind(this);
    this.addMessage = this.addMessage.bind(this);
    this.sendWSMessage = this.sendWSMessage.bind(this);
    this.connect = this.connect.bind(this);
    this.check = this.check.bind(this);
    this.search = this.search.bind(this);
  }

  handleSearchInput(searchInput) {
    this.setState({searchInput: searchInput});
  }

  handleSearchSubmit(event) {
    event.preventDefault();
    if (this.state.searchInput.length == 0) {
      return;
    }

    this.search(this.state.searchInput);
  }

  sendWSMessage(message) {
    // add messages to queue if connection is not ready
    if (!this.state.ws || this.state.ws.readyState != WebSocket.OPEN) {
      if (this.state.ws) {
        console.log("ws.readyState " + this.state.ws.readyState);
      }
      if (!this.ws_msgs) this.ws_msgs = []
      console.log("Added message " + message + " to queue");
      this.ws_msgs.push(message)
    } else {
      console.log("Sending message " + message);
      this.state.ws.send(message)
    }
  }

  search(input) {
    try {
      this.state.messages = [];
      this.sendWSMessage(JSON.stringify({
        'action': 'new',
        'message': input,
      }));
    } catch (error) {
      console.log(error);
    }
  }

  handleDeleteApp(event) {
    try {
      this.sendWSMessage(JSON.stringify({
        'action': 'delete',
        'message': this.state.appName
      }))
      // Remove message with app-label from the list
      let newMessages = this.state.messages.slice(1, this.state.messages.length)
      this.setState(state => ({ messages: newMessages, appName: null }))
    } catch (error) {
      console.log(error)
    }
  }

  addMessage(message) {
    this.setState(state => ({ messages: [...state.messages, message] }))
    if (message.action === "app-label") {
      this.setState(state => ({appName: message.message}))
    }
    if (message.action === "done") {
      // Remove message with progress from the list
      let newMessages = this.state.messages.filter(function(message){
        return message.action != "progress";
      });
      this.setState(state => ({ messages: newMessages }))
    }
    if (message.action === "rquota") {
      let rquotaStatus = JSON.parse(message.message)
      this.setState(state => ({
        resourceQuota: {
          used: rquotaStatus.used,
          hard: rquotaStatus.hard,
        }
      }))
    }
  }

  check () {
    const { ws } = this.state;
    if (!ws || ws.readyState == WebSocket.CLOSED) this.connect(); //check if websocket instance is closed, if so call `connect` function.
    };

  connect () {
    var loc = window.location;
    var ws_uri;
    if (loc.protocol === "https:") {
        ws_uri = "wss:";
    } else {
        ws_uri = "ws:";
    }
    ws_uri += "//" + loc.host;
    ws_uri += "/ws/status";
    var ws = new WebSocket(ws_uri);
    let that = this;
    var connectInterval;

    // websocket onopen event listener
    ws.onopen = () => {
      console.log("websocket connected");
      that.setState({ ws: ws });

      that.timeout = 250; // reset timer to 250 on open of websocket connection
      clearTimeout(connectInterval); // clear Interval on on open of websocket connection

      // Send messages if there's a queue
      while (that.ws_msgs && that.ws_msgs.length > 0) {
        ws.send(that.ws_msgs.pop())
      }
    };

    // websocket onclose event listener
    ws.onclose = e => {
      console.log(
        `Socket is closed. Reconnect will be attempted in ${Math.min(
            10000 / 1000,
            (that.timeout + that.timeout) / 1000
        )} second.`,
        e.reason
      );

      that.timeout = that.timeout + that.timeout; //increment retry interval
      connectInterval = setTimeout(this.check, Math.min(10000, that.timeout)); //call check function after timeout
    };

    // websocket onerror event listener
    ws.onerror = err => {
      console.error(
        "Socket encountered error: ",
        err.message,
        "Closing socket"
      );

      ws.close();
    };

    ws.onmessage = evt => {
      console.log("Received " + evt.data);
      const message = JSON.parse(evt.data);
      this.addMessage(message)
    }

    this.setState({ ws: ws });
  }

  componentDidMount() {
    this.check();
    this.timeout = 0;
    if (!this.state.searchInput) {
      let params = (new URL(window.location)).searchParams;
      let searchInput = params.get('search');
      if (searchInput && searchInput != this.state.querySearch) {
        this.state.querySearch = searchInput;
        this.handleSearchInput(searchInput);
        this.search(searchInput);
      }
    }
  }

  render() {
    let messages;
    let searchClass;
    if(this.state.appName != null) {
      messages =
        <Status messages={this.state.messages} />
      searchClass = null;
    } else {
      messages = [];
      searchClass = 'search-center';
    }

    return (
      <div className={searchClass}>
        <h3>PromeCIeus</h3>
        <SearchBar
          searchInput={this.state.searchInput}
          onSearchInput={this.handleSearchInput}
          onSearchSubmit={this.handleSearchSubmit}
          onDeleteApp={this.handleDeleteApp}
          appName={this.state.appName}
        />
        <ReactBootstrap.Row>
          <ReactBootstrap.Col xs={4}/>
          <ReactBootstrap.Col xs={4}>
            <ResourceQuotaStatus
              resourceQuota={this.state.resourceQuota || null}
            />
          </ReactBootstrap.Col>
          <ReactBootstrap.Col xs={4}/>
        </ReactBootstrap.Row>
        <br />
        {messages}
      </div>
    );
  }
}

ReactDOM.render(
  <SearchForm />,
  document.getElementById('container')
);
